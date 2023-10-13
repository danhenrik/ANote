package main

import (
	"anote/es_replicate/constants"
	"anote/es_replicate/database"
	"anote/es_replicate/es"
	"anote/es_replicate/types"
	"log"
)

// Elasticsearch

func init() {
	constants.Config()
	database.Connect()
	es.Connect()
}

func main() {
	updateChan := make(chan []types.Update)
	go database.ListenForUpdates(updateChan)
	for {
		updates := <-updateChan
		log.Println("Received updates:", updates)

		// write to elasticsearch
		es.ESIndex(updates)
		log.Println("Indexed updates:", updates)
	}
}
