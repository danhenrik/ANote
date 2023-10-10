package es

import (
	"anote/es_replicate/constants"
	db "anote/es_replicate/database"
	"anote/es_replicate/types"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/refresh"
)

var es *elasticsearch.TypedClient

func Connect() {
	cfg := elasticsearch.Config{
		Addresses: []string{fmt.Sprintf("%s:9200", constants.ES_ADDR)},
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

func ESIndex(updates []types.Update) {
	log.Println("Indexing to Elasticsearch...")
	indexName := "notes"

	for _, update := range updates {
		for _, change := range update.Change {
			if change.Kind == "delete" {
				switch change.Table {
				case "notes":
					noteId := change.OldKeys.KeyValues[0]
					log.Println("Deleting note w/ id", noteId, "from Elasticsearch")
					_, err := es.Delete(indexName, noteId).Do(context.Background())
					if err != nil {
						log.Println("Error deleting note w/ id", noteId, "from Elasticsearch: ", err)
					}
					log.Println("Successfully deleted")
					return
				case "likes":
					// remove like
				case "comments":
					// remove comment
				case "note_tags":
					// remove tag
				}
			}

			note := types.Note{}
			switch change.Table {
			case "notes":
				idIndex := strArr(change.ColumnNames).indexOf("id")
				noteId := change.ColumnValues[idIndex]

				// get note info
				res, err := es.Get(indexName, noteId).Do(context.Background())
				if err != nil {
					log.Println("Error on GET ES request:", err)
				}

				err = json.Unmarshal(res.Source_, &note)
				if err != nil {
					log.Println("Error on JSON decoding:", err)
				}

				// update or insert note info
				note.Id = change.ColumnValues[idIndex]
				titleIndex := strArr(change.ColumnNames).indexOf("title")
				note.Title = change.ColumnValues[titleIndex]
				contentIndex := strArr(change.ColumnNames).indexOf("content")
				note.Content = change.ColumnValues[contentIndex]
				publishedDateIndex := strArr(change.ColumnNames).indexOf("created_at")
				note.PublishedDate = change.ColumnValues[publishedDateIndex]
				updatedDateIndex := strArr(change.ColumnNames).indexOf("updated_at")
				note.UpdatedDate = change.ColumnValues[updatedDateIndex]
				authorIndex := strArr(change.ColumnNames).indexOf("user_id")
				note.Author = change.ColumnValues[authorIndex]

				// jsonString, err := json.Marshal(note)
				_, err = es.Index(indexName).
					Id(noteId).
					Request(note).
					Refresh(refresh.True).
					Do(context.Background())
				if err != nil {
					log.Println("Error on ES indexing:", err)
				}

				log.Println("Successfully indexed note w/ id", note.Id)
				db.SetLastItemReplicated(update.WalId)
				break
			case "likes":
				// Likes can only be inserted and deleted
				if change.Kind == "insert" {
					idIndex := strArr(change.ColumnNames).indexOf("note_id")
					noteId := change.ColumnValues[idIndex]

					// get note info
					res, err := es.Get(indexName, noteId).Do(context.Background())
					if err != nil {
						log.Println("Error on GET ES request:", err)
					}

					err = json.Unmarshal(res.Source_, &note)
					if err != nil {
						log.Println("Error on JSON decoding:", err)
					}

					// update or insert note info
					idIndex = strArr(change.ColumnNames).indexOf("id")
					userIdIndex := strArr(change.ColumnNames).indexOf("user_id")
					like := types.Like{Id: change.ColumnValues[idIndex], UserId: change.ColumnValues[userIdIndex]}
					note.Likes = append(note.Likes, like)
					note.LikesCount = len(note.Likes)

					// jsonString, err := json.Marshal(note)
					_, err = es.
						Index(indexName).
						Id(noteId).
						Request(note).
						Do(context.Background())
					if err != nil {
						log.Println("Error on ES indexing:", err)
					}

					log.Println("Successfully indexed likes from note w/ id", note.Id)
					db.SetLastItemReplicated(update.WalId)
					break
				}
				log.Println("Likes can't be updated")
			case "comments":
				// comments can only be inserted and deleted
				if change.Kind == "insert" {
					idIndex := strArr(change.ColumnNames).indexOf("note_id")
					noteId := change.ColumnValues[idIndex]

					// get note info
					res, err := es.Get(indexName, noteId).Do(context.Background())
					if err != nil {
						log.Println("Error on GET ES request:", err)
					}
					err = json.Unmarshal(res.Source_, &note)
					if err != nil {
						log.Println("Error on JSON decoding:", err)
					}

					// update or insert note info
					idIndex = strArr(change.ColumnNames).indexOf("id")
					userIdIndex := strArr(change.ColumnNames).indexOf("user_id")
					comment := types.Comment{Id: change.ColumnValues[idIndex], UserId: change.ColumnValues[userIdIndex]}

					note.Commenters = append(note.Commenters, comment)
					note.CommentCount = len(note.Commenters)

					// jsonString, err := json.Marshal(note)
					_, err = es.
						Index(indexName).
						Id(noteId).
						Request(note).
						Do(context.Background())
					if err != nil {
						log.Println("Error on ES indexing:", err)
					}

					log.Println("Successfully indexed comments from note w/ id", note.Id)
					db.SetLastItemReplicated(update.WalId)
					break
				}
				log.Println("Comments can't be updated")
			case "note_tags":
				// tags can only be inserted and deleted
				if change.Kind == "insert" {
					idIndex := strArr(change.ColumnNames).indexOf("note_id")
					noteId := change.ColumnValues[idIndex]

					// get note info
					res, err := es.Get(indexName, noteId).Do(context.Background())
					if err != nil {
						log.Println("Error on GET ES request:", err)
					}
					err = json.Unmarshal(res.Source_, &note)
					if err != nil {
						log.Println("Error on JSON decoding:", err)
					}

					// update or insert note info
					idIndex = strArr(change.ColumnNames).indexOf("id")
					tagIdIndex := strArr(change.ColumnNames).indexOf("tag_id")
					tag := db.GetTagById(change.ColumnValues[tagIdIndex])
					tag.Id = change.ColumnValues[idIndex]

					note.Tags = append(note.Tags, tag)

					// jsonString, err := json.Marshal(note)
					_, err = es.
						Index(indexName).
						Id(noteId).
						Request(note).
						Do(context.Background())

					if err != nil {
						log.Println("Error on ES indexing:", err)
					}

					db.SetLastItemReplicated(update.WalId)
					break
				}
				log.Println("Tags can't be updated")
			}
		}
	}
}
