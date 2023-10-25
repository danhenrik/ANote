package es

import (
	"anote/internal/constants"
	"anote/internal/errors"
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

type ESClient struct {
	client *elasticsearch.TypedClient
}

func NewESClient() *ESClient {
	cfg := elasticsearch.Config{
		Addresses: []string{fmt.Sprintf("http://%s:9200", constants.ES_ADDR)},
	}

	client, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		log.Fatalf("[NewESClient] Error connecting to Elasticsearch: %s", err)
	}
	log.Println("[NewESClient] Connected to Elasticsearch")
	return &ESClient{client: client}
}

func (this *ESClient) Search(index, query string) ([]any, *errors.AppError) {
	res, err := this.client.
		Search().
		Index(index).
		Raw(strings.NewReader(query)).
		Do(context.Background())
	if err != nil {
		log.Printf("[Query] Error getting response: %s", err)
		return nil, errors.NewAppError(500, err.Error())
	}

	var result []any
	for _, hit := range res.Hits.Hits {
		result = append(result, hit.Source_)
	}
	return result, nil
}
