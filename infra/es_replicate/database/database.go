package database

import (
	"anote/es_replicate/constants"
	"anote/es_replicate/types"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/lib/pq"
)

var db *sql.DB
var connStr string

func Connect() {
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		constants.DB_ADDR,
		constants.DB_USER,
		constants.DB_PWD,
		constants.DB_NAME,
	)

	database, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to PostgreSQL: %s", err)
	}
	db = database
	log.Println("Connected to PostgreSQL")
}

type UpdateRow struct {
	lsn  []uint8
	xid  int64
	data string
}

func GetUpdateFromPG() (updatesList []types.Update) {
	updates, err := db.Query("SELECT * FROM pg_logical_slot_peek_changes('es_replication_slot', NULL, NULL);")
	if err != nil {
		log.Println("Error on query:", err)
	}
	defer updates.Close()

	updatesList = []types.Update{}
	for updates.Next() {
		row := UpdateRow{}
		err := updates.Scan(&row.lsn, &row.xid, &row.data)
		if err != nil {
			log.Println("Error on query scan:", err)
		}

		// note: Could probably have a more performant encoding to data changes
		update := types.Update{}
		err = json.Unmarshal([]byte(row.data), &update)
		if err != nil {
			log.Println("Error on JSON Unmarshal: ", err)
		}
		update.WalId = row.lsn
		updatesList = append(updatesList, update)
	}
	return updatesList
}

func ListenForUpdates(updateChan chan<- []types.Update) {
	// create a new listener
	listener := pq.NewListener(connStr, 10*time.Second, time.Minute, nil)
	defer listener.Close()
	err := listener.Listen("es_replicate")
	if err != nil {
		log.Fatalln("Error on listen:", err)
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
		case <-time.After(5 * time.Minute):
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
	_, err := db.Exec("SELECT pg_replication_slot_advance('es_replication_slot', $1);", lsn)
	if err != nil {
		log.Println("Error on replication callback:", err)
	}
}

func GetTagById(id string) (tag types.Tag) {
	query := db.QueryRow("SELECT * FROM tags WHERE id = $1;", id)
	err := query.Scan(&tag.Id, &tag.Name)
	if err != nil {
		log.Println("Error on query scan:", err)
	}
	return tag
}
