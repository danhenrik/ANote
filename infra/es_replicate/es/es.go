package es

import (
	"anote/es_replicate/constants"
	db "anote/es_replicate/database"
	appTypes "anote/es_replicate/types"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	updateES "github.com/elastic/go-elasticsearch/v8/typedapi/core/update"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/refresh"
)

var es *elasticsearch.TypedClient

func Connect() {
	cfg := elasticsearch.Config{
		Addresses: []string{fmt.Sprintf("http://%s:9200", constants.ES_ADDR)},
	}

	client, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		log.Fatalf("Error connecting to Elasticsearch: %s", err)
	}
	es = client
	log.Println("Connected to Elasticsearch")
}

type strArr []string

func (arr strArr) indexOf(element string) int {
	for i, v := range arr {
		if v == element {
			return i
		}
	}
	return -1
}

const indexName = "notes"

func GetNote(change *appTypes.Change) *appTypes.Note {
	idIndex := strArr(change.ColumnNames).indexOf("note_id")
	noteId := change.ColumnValues[idIndex]

	// get note info
	res, err := es.Get(indexName, noteId).Do(context.Background())
	if err != nil {
		log.Println("Error on GET ES request:", err)
		return nil
	}

	if !res.Found {
		log.Println("Note w/ id", noteId, "not found in Elasticsearch")
		return nil
	}

	note := appTypes.Note{}
	if err = json.Unmarshal(res.Source_, &note); err != nil {
		log.Println("Error on JSON decoding:", err)
		return nil
	}
	return &note
}

func IndexNote(note *appTypes.Note) bool {
	_, err := es.Index(indexName).
		Id(note.Id).
		Request(note).
		Refresh(refresh.True).
		Do(context.Background())
	if err != nil {
		log.Println("Error on ES indexing:", err)
		return false
	}
	return true
}

func ESIndex(updates []appTypes.Update) {
	log.Println("Indexing to Elasticsearch...")
	for _, update := range updates {
		for _, change := range update.Change {
			if change.Kind == "delete" {
				switch change.Table {
				case "notes":
					idIdx := strArr(change.OldKeys.KeyNames).indexOf("id")
					noteId := change.OldKeys.KeyValues[idIdx]
					log.Println("Deleting note w/ id", noteId, "from Elasticsearch")
					_, err := es.Delete(indexName, noteId).Do(context.Background())
					if err != nil {
						log.Println("Error deleting note w/ id", noteId, "from Elasticsearch: ", err)
						return
					}
					log.Println("Successfully deleted note")
					db.SetLastItemReplicated(update.WalId)
					break
				case "likes":
					idIdx := strArr(change.OldKeys.KeyNames).indexOf("id")
					likeId := change.OldKeys.KeyValues[idIdx]
					log.Println("Deleting like w/ id", likeId, "from Elasticsearch")
					res, err := es.Search().
						Index(indexName).
						Query(&types.Query{
							Match: map[string]types.MatchQuery{
								"like.id": {Query: likeId},
							},
						}).
						Do(context.Background())
					if err != nil {
						log.Println("Error finding note to delete with like w/ id", likeId, "from Elasticsearch: ", err)
						return
					}
					if len(res.Hits.Hits) == 0 {
						log.Println("Note not found, can't delete like")
						return
					}
					var note appTypes.Note
					if err = json.Unmarshal(res.Hits.Hits[0].Source_, &note); err != nil {
						log.Println("Error on JSON decoding:", err)
						return
					}
					for idx, like := range note.Likes {
						if like.Id == likeId {
							note.Likes = append(note.Likes[:idx], note.Likes[idx+1:]...)
							break
						}
					}
					note.LikesCount = len(note.Likes)

					jsonString, err := json.Marshal(note)
					if err != nil {
						log.Println("Error on JSON encoding before update:", err)
						return
					}
					_, err = es.Update(indexName, note.Id).
						Request(&updateES.Request{Doc: jsonString}).
						Refresh(refresh.True).
						Do(context.Background())
					if err != nil {
						log.Println("Error updating note w/ id", note.Id, "after deleting like w/ id", likeId, "from Elasticsearch: ", err)
						return
					}
					log.Println("Successfully deleted like")
					db.SetLastItemReplicated(update.WalId)
					break
				case "comments":
					idIdx := strArr(change.OldKeys.KeyNames).indexOf("id")
					commentId := change.OldKeys.KeyValues[idIdx]
					log.Println("Deleting comment w/ id", commentId, "from Elasticsearch")
					res, err := es.Search().
						Index(indexName).
						Query(&types.Query{
							Match: map[string]types.MatchQuery{
								"commenters.id": {Query: commentId},
							},
						}).
						Do(context.Background())
					if err != nil {
						log.Println("Error finding note to delete with comment w/ id", commentId, "from Elasticsearch: ", err)
						return
					}
					if len(res.Hits.Hits) == 0 {
						log.Println("Note not found, can't delete comment")
						return
					}
					var note appTypes.Note
					if err = json.Unmarshal(res.Hits.Hits[0].Source_, &note); err != nil {
						log.Println("Error on JSON decoding:", err)
						return
					}
					for idx, comment := range note.Commenters {
						if comment.Id == commentId {
							note.Commenters = append(note.Commenters[:idx], note.Commenters[idx+1:]...)
							break
						}
					}
					note.CommentCount = len(note.Commenters)

					jsonString, err := json.Marshal(note)
					if err != nil {
						log.Println("Error on JSON encoding before update:", err)
						return
					}
					_, err = es.Update(indexName, note.Id).
						Request(&updateES.Request{Doc: jsonString}).
						Refresh(refresh.True).
						Do(context.Background())
					if err != nil {
						log.Println("Error updating note w/ id", note.Id, "after deleting comment w/ id", commentId, "from Elasticsearch: ", err)
						return
					}
					log.Println("Successfully deleted like")
					db.SetLastItemReplicated(update.WalId)
					break
				case "note_tags":
					idIdx := strArr(change.OldKeys.KeyNames).indexOf("id")
					tagId := change.OldKeys.KeyValues[idIdx]
					log.Println("Deleting tag w/ id", tagId, "from Elasticsearch")
					res, err := es.Search().
						Index(indexName).
						Query(&types.Query{
							Match: map[string]types.MatchQuery{
								"tags.r_id": {Query: tagId},
							},
						}).
						Do(context.Background())
					if err != nil {
						log.Println("Error finding note to delete with tag w/ id", tagId, "from Elasticsearch: ", err)
						return
					}
					if len(res.Hits.Hits) == 0 {
						log.Println("Note not found, can't delete tag")
						return
					}
					var note appTypes.Note
					if err = json.Unmarshal(res.Hits.Hits[0].Source_, &note); err != nil {
						log.Println("Error on JSON decoding:", err)
						return
					}
					for idx, tag := range note.Tags {
						if tag.RId == tagId {
							note.Tags = append(note.Tags[:idx], note.Tags[idx+1:]...)
							break
						}
					}

					jsonString, err := json.Marshal(note)
					if err != nil {
						log.Println("Error on JSON encoding before update:", err)
						return
					}
					_, err = es.Update(indexName, note.Id).
						Request(&updateES.Request{Doc: jsonString}).
						Refresh(refresh.True).
						Do(context.Background())
					if err != nil {
						log.Println("Error updating note w/ id", note.Id, "after deleting tag w/ id", tagId, "from Elasticsearch: ", err)
						return
					}
					log.Println("Successfully deleted tag")
					db.SetLastItemReplicated(update.WalId)
					break
				case "community_notes":
					idIdx := strArr(change.OldKeys.KeyNames).indexOf("id")
					communityId := change.OldKeys.KeyValues[idIdx]
					log.Println("Deleting community w/ id", communityId, "from Elasticsearch")
					res, err := es.Search().
						Index(indexName).
						Query(&types.Query{
							Match: map[string]types.MatchQuery{
								"communities.r_id": {Query: communityId},
							},
						}).
						Do(context.Background())
					if err != nil {
						log.Println("Error finding note to delete with community w/ id", communityId, "from Elasticsearch: ", err)
						return
					}
					if len(res.Hits.Hits) == 0 {
						log.Println("Note not found, can't delete community")
						return
					}
					var note appTypes.Note
					if err = json.Unmarshal(res.Hits.Hits[0].Source_, &note); err != nil {
						log.Println("Error on JSON decoding:", err)
						return
					}
					for idx, community := range note.Communities {
						if community.RId == communityId {
							note.Communities = append(note.Communities[:idx], note.Communities[idx+1:]...)
							break
						}
					}

					jsonString, err := json.Marshal(note)
					if err != nil {
						log.Println("Error on JSON encoding before update:", err)
						return
					}
					_, err = es.Update(indexName, note.Id).
						Request(&updateES.Request{Doc: jsonString}).
						Refresh(refresh.True).
						Do(context.Background())
					if err != nil {
						log.Println("Error updating note w/ id", note.Id, "after deleting community w/ id", communityId, "from Elasticsearch: ", err)
						return
					}
					log.Println("Successfully deleted tag")
					db.SetLastItemReplicated(update.WalId)
					break
				}
			}

			switch change.Table {
			case "notes":
				idIdx := strArr(change.ColumnNames).indexOf("id")
				noteId := change.ColumnValues[idIdx]

				// get note info
				res, err := es.Get(indexName, noteId).Do(context.Background())
				if err != nil {
					log.Println("Error on GET ES request:", err)
					return
				}

				note := appTypes.Note{}
				if res.Found {
					if err = json.Unmarshal(res.Source_, &note); err != nil {
						log.Println("Error on JSON decoding:", err)
						return
					}
					log.Println("Note w/ id", noteId, "found in Elasticsearch, updating...")
				} else {
					log.Println("Note w/ id", noteId, "not found in Elasticsearch, inserting...")
				}

				// insert or update note info
				note.Id = noteId
				titleIdx := strArr(change.ColumnNames).indexOf("title")
				note.Title = change.ColumnValues[titleIdx]
				contentIdx := strArr(change.ColumnNames).indexOf("content")
				note.Content = change.ColumnValues[contentIdx]
				publishedDateIdx := strArr(change.ColumnNames).indexOf("created_at")
				note.PublishedDate = change.ColumnValues[publishedDateIdx]
				updatedDateIdx := strArr(change.ColumnNames).indexOf("updated_at")
				note.UpdatedDate = change.ColumnValues[updatedDateIdx]
				authorIdx := strArr(change.ColumnNames).indexOf("author_id")
				note.Author = change.ColumnValues[authorIdx]

				indexed := IndexNote(&note)
				if !indexed {
					log.Println("Unable to index note w/ id", note.Id)
					return
				}
				log.Println("Successfully indexed note w/ id", note.Id)
				db.SetLastItemReplicated(update.WalId)
				break
			case "likes":
				// Likes can only be inserted and deleted
				if change.Kind == "insert" {
					note := GetNote(&change)
					if note == nil {
						log.Println("Note not found, can't insert like")
						return
					}

					// insert likes
					idIdx := strArr(change.ColumnNames).indexOf("id")
					userIdIdx := strArr(change.ColumnNames).indexOf("user_id")
					like := appTypes.Like{Id: change.ColumnValues[idIdx], UserId: change.ColumnValues[userIdIdx]}
					note.Likes = append(note.Likes, like)
					note.LikesCount = len(note.Likes)

					indexed := IndexNote(note)
					if !indexed {
						log.Println("Unable to index like", like, "for note w/ id", note.Id)
						return
					}
					log.Println("Successfully indexed likes from note w/ id", note.Id)
					db.SetLastItemReplicated(update.WalId)
					break
				}
				log.Println("Likes can't be updated")
			case "comments":
				// comments can only be inserted and deleted
				if change.Kind == "insert" {
					note := GetNote(&change)
					if note == nil {
						log.Println("Note not found, can't insert comment")
						return
					}

					// insert comment
					idIdx := strArr(change.ColumnNames).indexOf("id")
					userIdIdx := strArr(change.ColumnNames).indexOf("user_id")
					comment := appTypes.Comment{Id: change.ColumnValues[idIdx], UserId: change.ColumnValues[userIdIdx]}
					note.Commenters = append(note.Commenters, comment)
					note.CommentCount = len(note.Commenters)

					indexed := IndexNote(note)
					if !indexed {
						log.Println("Unable to index comment", comment, "for note w/ id", note.Id)
						return
					}
					log.Println("Successfully indexed comments from note w/ id", note.Id)
					db.SetLastItemReplicated(update.WalId)
					break
				}
				log.Println("Comments can't be updated")
			case "note_tags":
				// tags can only be inserted and deleted
				if change.Kind == "insert" {
					note := GetNote(&change)
					if note == nil {
						log.Println("Note not found, can't insert tag")
						return
					}

					// insert tag
					tagIdIdx := strArr(change.ColumnNames).indexOf("tag_id")
					tagId := change.ColumnValues[tagIdIdx]
					log.Println("tagId:", tagId)
					tag := db.GetTagById(tagId)
					idIndex := strArr(change.ColumnNames).indexOf("id")
					tag.RId = change.ColumnValues[idIndex]
					note.Tags = append(note.Tags, *tag)

					indexed := IndexNote(note)
					if !indexed {
						log.Println("Unable to index tag", tag, "for note w/ id", note.Id)
						return
					}
					log.Println("Successfully indexed tags from note w/ id", note.Id)
					db.SetLastItemReplicated(update.WalId)
					break
				}
				log.Println("Tags can't be updated")
			case "community_notes":
				// community_notes can only be inserted and deleted
				if change.Kind == "insert" {
					note := GetNote(&change)
					if note == nil {
						log.Println("Note not found, can't insert community note")
						return
					}

					// insert community
					communityIdIdx := strArr(change.ColumnNames).indexOf("community_id")
					community := db.GetCommunityById(change.ColumnValues[communityIdIdx])

					idIdx := strArr(change.ColumnNames).indexOf("id")
					community.RId = change.ColumnValues[idIdx]
					note.Communities = append(note.Communities, *community)

					indexed := IndexNote(note)
					if !indexed {
						log.Println("Unable to index community", community, "for note w/ id", note.Id)
						return
					}
					log.Println("Successfully indexed community notes from note w/ id", note.Id)
					db.SetLastItemReplicated(update.WalId)
					break
				}
				log.Println("Community notes can't be updated")
			default:
				log.Println("Table", change.Table, "not supported")
				db.SetLastItemReplicated(update.WalId)
			}
		}
	}
}
