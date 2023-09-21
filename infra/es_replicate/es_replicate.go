package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
)

type UpdateRow struct {
	lsn  []uint8
	xid  int64
	data string
}

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

type SharedContainer struct {
	connStr string
	db      *sql.DB
	es      *elasticsearch.Client
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

var container SharedContainer

func GetUpdateFromPG() (updatesList []Update) {
	db := container.db
	updates, err := db.Query("SELECT * FROM pg_logical_slot_peek_changes('es_replication_slot', NULL, NULL);")
	if err != nil {
		fmt.Println("Error on query: ", err)
	}
	defer updates.Close()

	updatesList = []Update{}
	for updates.Next() {
		row := UpdateRow{}
		err := updates.Scan(&row.lsn, &row.xid, &row.data)
		if err != nil {
			fmt.Print("Error on query scan: ", err)
		}

		// note: Could probably have a more performant encoding to data changes
		update := Update{}
		err = json.Unmarshal([]byte(row.data), &update)
		if err != nil {
			fmt.Print("Error on JSON Unmarshal: ", err)
		}
		update.WalId = row.lsn
		updatesList = append(updatesList, update)
	}
	return updatesList
}

func ListenForUpdates(updateChan chan<- []Update) {
	// create a new listener
	connStr := container.connStr
	listener := pq.NewListener(connStr, 10*time.Second, time.Minute, nil)
	defer listener.Close()
	err := listener.Listen("es_replicate")
	if err != nil {
		fmt.Println("Error on listen: ", err)
	} else {
		fmt.Println("Listening for notifications...")
	}

	// continuously listen for notifications
	timeIdle := 0
	for {
		select {
		case <-listener.Notify:
			fmt.Println("Received update notification, reaching for changes...")
			updatesList := GetUpdateFromPG()
			if len(updatesList) > 0 {
				fmt.Println("Received ", len(updatesList), " updates")
				fmt.Println(updatesList)
				updateChan <- updatesList
				timeIdle = 0
			} else {
				fmt.Println("Notification sent falsely")
			}
		case <-time.After(time.Minute):
			timeIdle++
			if timeIdle == 0 {
				fmt.Println("Received no events for ", timeIdle, " minute")
			} else {
				fmt.Println("Received no events for ", timeIdle, " minutes")
			}
			updatesList := GetUpdateFromPG()
			if len(updatesList) > 0 {
				fmt.Println("Found ", len(updatesList), " updates!")
				fmt.Println(updatesList)
				updateChan <- updatesList
				timeIdle = 0
			}
		}
	}
}

type Note struct {
	Id            string   `json:"id"`
	Title         string   `json:"title"`
	Content       string   `json:"content"`
	LikesCount    int      `json:"likes_count"`
	Likes         []string `json:"likes"` // usernames (id)
	PublishedDate string   `json:"published_date"`
	UpdatedDate   string   `json:"updated_date"`
	Author        string   `json:"author"` // username (id)
	Tags          []string `json:"tags"`   // tag names (id)
	CommentCount  int      `json:"comment_count"`
	Commenters    []string `json:"commenters"` // usernames (id)
}

func ESIndex(updates []Update) {
	fmt.Println("Indexing to Elasticsearch...")
	es := container.es
	indexName := "notes"

	for _, update := range updates {
		for _, change := range update.Change {
			if change.Kind == "delete" {
				switch change.Table {
				case "notes":
					noteId := change.OldKeys.KeyValues[0]
					log.Println("Deleting note w/ id ", noteId, " from Elasticsearch")
					_, err := es.Delete(indexName, noteId)
					if err != nil {
						log.Println("Error deleting note w/ id, ", noteId, " from Elasticsearch: ", err)
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
				log.Println(change)
				noteId := change.OldKeys.KeyValues[0]
				res, err := es.Get(indexName, noteId)
				// get note info
				if err != nil {
					fmt.Print("Error on GET ES request: ", err)
				}

				json.NewDecoder(res.Body).Decode(&note)
				log.Println(note)
				if err != nil {
					fmt.Print("Error on JSON decoding: ", err)
				}

				// update or insert note info
				titleIndex := strArr(change.ColumnNames).indexOf("title")
				note.Title = change.ColumnValues[titleIndex]
				contentIndex := strArr(change.ColumnNames).indexOf("content")
				note.Content = change.ColumnValues[contentIndex]
				updatedDateIndex := strArr(change.ColumnNames).indexOf("updated_at")
				note.UpdatedDate = change.ColumnValues[updatedDateIndex]

				if change.Kind == "insert" {
					idIndex := strArr(change.ColumnNames).indexOf("id")
					note.Id = change.ColumnValues[idIndex]
					authorIndex := strArr(change.ColumnNames).indexOf("user_id")
					note.Author = change.ColumnValues[authorIndex]
					publishedDateIndex := strArr(change.ColumnNames).indexOf("created_at")
					note.PublishedDate = change.ColumnValues[publishedDateIndex]
				}

				jsonString, err := json.Marshal(note)
				res, err = es.Index(
					indexName,
					strings.NewReader(string(jsonString)),
					es.Index.WithWaitForActiveShards("all"),
				)
				if err != nil {
					fmt.Println("Error on ES indexing: ", err)
				}
				log.Println("Successfully indexed note w/ id ", note.Id, " to Elasticsearch")

			case "likes":
				// 0: ID, 1: user_id, 2: note_id
				noteId := change.ColumnValues[2]
				res, err := es.Get(indexName, noteId)
				if err != nil {
					fmt.Print("Error on GET ES request: ", err)
				}

				json.NewDecoder(res.Body).Decode(&note)
				log.Println(note)
				if err != nil {
					fmt.Print("Error on JSON decoding: ", err)
				}
			case "comments":
				// 0: ID, 1: user_id, 2: note_id, 3: content

			case "note_tags":
			}
		}

		// esIndex := "notes"
		es.Info()
		// es.Index(esIndex)
	}
}

func main() {
	// Connect to PostgreSQL
	godotenv.Load("../../.env")
	var pg_user string = os.Getenv("PG_USER")
	var pg_password string = os.Getenv("PG_PASSWORD")
	var pg_dbname string = os.Getenv("PG_DATABASE")
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", pg_user, pg_password, pg_dbname)
	container.connStr = connStr

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to PostgreSQL: %s", err)
	}
	log.Println("Connected to PostgreSQL")
	container.db = db
	defer db.Close()

	// Connect to Elasticsearch
	cfg := elasticsearch.Config{
		Addresses: []string{"http://elasticsearch:9200"},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error connecting to Elasticsearch: %s", err)
	}
	log.Println("Connected to Elasticsearch")
	container.es = es

	updateChan := make(chan []Update)
	go ListenForUpdates(updateChan)
	for {
		updates := <-updateChan

		// write to elasticsearch
		ESIndex(updates)
	}
}
