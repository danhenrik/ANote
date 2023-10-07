package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/refresh"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
)

// Shared
type SharedContainer struct {
	connStr string
	db      *sql.DB
	es      *elasticsearch.TypedClient
}

var Container SharedContainer

type Update struct {
	WalId  []uint8
	Change []struct {
		Kind         string   `json:"kind"`
		Schema       string   `json:"schema"`
		Table        string   `json:"table"`
		ColumnNames  []string `json:"columnnames"`
		ColumnTypes  []string `json:"columntypes"`
		ColumnValues []string `json:"columnvalues"`
		OldKeys      struct {
			KeyNames  []string `json:"keynames"`
			KeyTypes  []string `json:"keytypes"`
			KeyValues []string `json:"keyvalues"`
		} `json:"oldkeys"`
	} `json:"change"`
}

// Postgres
type UpdateRow struct {
	lsn  []uint8
	xid  int64
	data string
}

func GetUpdateFromPG() (updatesList []Update) {
	db := Container.db
	updates, err := db.Query("SELECT * FROM pg_logical_slot_peek_changes('es_replication_slot', NULL, NULL);")
	if err != nil {
		log.Println("Error on query:", err)
	}
	defer updates.Close()

	updatesList = []Update{}
	for updates.Next() {
		row := UpdateRow{}
		err := updates.Scan(&row.lsn, &row.xid, &row.data)
		if err != nil {
			log.Println("Error on query scan:", err)
		}

		// note: Could probably have a more performant encoding to data changes
		update := Update{}
		err = json.Unmarshal([]byte(row.data), &update)
		if err != nil {
			log.Println("Error on JSON Unmarshal: ", err)
		}
		update.WalId = row.lsn
		updatesList = append(updatesList, update)
	}
	return updatesList
}

func ListenForUpdates(updateChan chan<- []Update) {
	// create a new listener
	connStr := Container.connStr
	listener := pq.NewListener(connStr, 10*time.Second, time.Minute, nil)
	defer listener.Close()
	err := listener.Listen("es_replicate")
	if err != nil {
		log.Println("Error on listen:", err)
	} else {
		log.Println("Listening for notifications...")
	}

	// continuously listen for notifications
	timeIdle := 0
	for {
		select {
		case <-listener.Notify:
			log.Println("Received update notification, reaching for changes...")
			updatesList := GetUpdateFromPG()
			if len(updatesList) > 0 {
				log.Println("Received", len(updatesList), "updates")
				log.Println(updatesList)
				updateChan <- updatesList
				timeIdle = 0
			} else {
				log.Println("Notification sent falsely")
			}
		case <-time.After(time.Minute):
			timeIdle++
			if timeIdle == 0 {
				log.Println("Received no events for", timeIdle, "minute")
			} else {
				log.Println("Received no events for", timeIdle, "minutes")
			}
			updatesList := GetUpdateFromPG()
			if len(updatesList) > 0 {
				log.Println("Found", len(updatesList), "updates!")
				log.Println(updatesList)
				updateChan <- updatesList
				timeIdle = 0
			}
		}
	}
}

func SetLastItemReplicated(lsn []uint8) {
	db := Container.db
	_, err := db.Exec("SELECT pg_replication_slot_advance('es_replication_slot', $1);", lsn)
	if err != nil {
		log.Println("Error on replication callback:", err)
	}
}

func GetTagById(id string) (tag Tag) {
	db := Container.db
	query := db.QueryRow("SELECT * FROM tags WHERE id = $1;", id)
	err := query.Scan(&tag.Id, &tag.Name)
	if err != nil {
		log.Println("Error on query scan:", err)
	}
	return tag
}

// Elasticsearch

type strArr []string

func (arr strArr) indexOf(element string) int {
	for i, v := range arr {
		if v == element {
			return i
		}
	}
	return -1
}

type Community struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Tag struct {
	Id    string `json:"id"`
	TagId string `json:"tag_id"`
	Name  string `json:"name"`
}

type Like struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
}

type Comment struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
}

type Note struct {
	Id            string      `json:"id"`
	Title         string      `json:"title"`
	Content       string      `json:"content"`
	PublishedDate string      `json:"published_date"`
	UpdatedDate   string      `json:"updated_date"`
	Author        string      `json:"author"` // username (id)
	Communities   []Community `json:"communities"`
	Tags          []Tag       `json:"tags"`
	LikesCount    int         `json:"likes_count"`
	Likes         []Like      `json:"likes"` // usernames (id)
	CommentCount  int         `json:"comment_count"`
	Commenters    []Comment   `json:"commenters"` // usernames (id)
}

func ESIndex(updates []Update) {
	log.Println("Indexing to Elasticsearch...")

	es := Container.es
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

			note := Note{}
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
				SetLastItemReplicated(update.WalId)
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
					like := Like{Id: change.ColumnValues[idIndex], UserId: change.ColumnValues[userIdIndex]}
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
					SetLastItemReplicated(update.WalId)
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
					comment := Comment{Id: change.ColumnValues[idIndex], UserId: change.ColumnValues[userIdIndex]}

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
					SetLastItemReplicated(update.WalId)
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
					tag := GetTagById(change.ColumnValues[tagIdIndex])
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

					SetLastItemReplicated(update.WalId)
					break
				}
				log.Println("Tags can't be updated")
			}
		}
	}
}

func main() {
	// Connect to PostgreSQL
	godotenv.Load("../../.env")
	var pg_user string = os.Getenv("PG_USER")
	var pg_password string = os.Getenv("PG_PASSWORD")
	var pg_dbname string = os.Getenv("PG_DATABASE")
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", pg_user, pg_password, pg_dbname)
	Container.connStr = connStr

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to PostgreSQL: %s", err)
	}
	log.Println("Connected to PostgreSQL")
	Container.db = db
	defer db.Close()

	// Connect to Elasticsearch
	// cfg := elasticsearch.Config{
	// 	Addresses: []string{"http://elasticsearch:9200"},
	// }
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}

	es, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		log.Fatalf("Error connecting to Elasticsearch: %s", err)
	}
	log.Println("Connected to Elasticsearch")
	Container.es = es

	updateChan := make(chan []Update)
	go ListenForUpdates(updateChan)
	for {
		updates := <-updateChan
		log.Println("Received updates:", updates)

		// write to elasticsearch
		ESIndex(updates)
		log.Println("Indexed updates:", updates)
	}
}
