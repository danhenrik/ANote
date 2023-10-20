package es

import (
	"anote/internal/constants"
	"context"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type QueryBuilder struct {
	index  string
	client *elasticsearch.TypedClient
	query  *types.Query
}

func NewQueryBuilder(indexName string) *QueryBuilder {
	cfg := elasticsearch.Config{
		Addresses: []string{fmt.Sprintf("http://%s:9200", constants.ES_ADDR)},
	}

	client, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		log.Fatalf("[NewQueryBuilder] Error connecting to Elasticsearch: %s", err)
	}
	log.Println("[NewQueryBuilder] Connected to Elasticsearch")

	return &QueryBuilder{
		index:  indexName,
		client: client,
		query: &types.Query{
			Match: map[string]types.MatchQuery{},
		},
	}
}

func (qb *QueryBuilder) AddQuery(field string, value string) *QueryBuilder {
	qb.query.Match[field] = types.MatchQuery{Query: value}
	return qb
}

func (qb *QueryBuilder) Query() any {
	res, err := qb.client.
		Search().
		Index(qb.index).
		Query(qb.query).
		Do(context.Background())

	if err != nil {
		log.Printf("[Query] Error getting response: %s", err)
		return err
	}
	return res
}
