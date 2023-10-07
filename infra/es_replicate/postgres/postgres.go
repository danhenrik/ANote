package postgres

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
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

var Container SharedContainer

func GetUpdateFromPG() (updatesList []Update) {
	db := Container.db
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
	connStr := Container.connStr
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
			// if len(updatesList) > 0 {
			fmt.Println("Received ", len(updatesList), " updates")
			fmt.Println(updatesList)
			updateChan <- updatesList
			timeIdle = 0
			// } else {
			// fmt.Println("Notification sent falsely")
			// }
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

func SetLastItemReplicated(lsn []uint8) {
	db := Container.db
	_, err := db.Exec("SELECT pg_replication_slot_advance('es_replication_slot', %u);", lsn)
	if err != nil {
		fmt.Println("Error on replication callback: ", err)
	}
}
