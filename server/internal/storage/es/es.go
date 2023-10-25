package es

import (
	"anote/internal/constants"
	"anote/internal/errors"
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type QueryBuilder struct {
	index      string
	client     *elasticsearch.TypedClient
	query      *strings.Builder
	sort       *strings.Builder
	queryCount int
	sortCount  int
	size       int
	page       int
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

	query := strings.Builder{}
	sort := strings.Builder{}

	query.WriteString(`{"query":{"bool":{"should": [{"bool":{"must":[`)
	sort.WriteString(`"sort":[`)
	return &QueryBuilder{
		index:      indexName,
		client:     client,
		query:      &query,
		sort:       &sort,
		queryCount: 0,
		sortCount:  0,
		size:       0,
		page:       0,
	}
}

func (qb *QueryBuilder) AddMatchQuery(field string, value string) *QueryBuilder {
	if qb.queryCount > 0 {
		qb.query.WriteString(",")
	}
	qb.query.WriteString(fmt.Sprintf(`{"match":{"%s":"%s"}}`, field, value))
	qb.queryCount++
	return qb
}

func (qb *QueryBuilder) AddIncludeQuery(field string, value []string) *QueryBuilder {
	if qb.queryCount > 0 {
		qb.query.WriteString(",")
	}
	termsStr := strings.Join(value, "\", \"")
	qb.query.WriteString(fmt.Sprintf(`{"terms":{"%s": ["%s"]}}`, field, termsStr))
	qb.queryCount++
	return qb
}

func (qb *QueryBuilder) AddWildcardQuery(field string, wildcard string) *QueryBuilder {
	if qb.queryCount > 0 {
		qb.query.WriteString(",")
	}
	qb.query.WriteString(fmt.Sprintf(`{"wildcard":{"%s":{"value": "%s"}}}`, field, wildcard))
	qb.queryCount++
	return qb
}

func (qb *QueryBuilder) AddRangeQuery(field string, lower string, upper string) *QueryBuilder {
	if qb.queryCount > 0 {
		qb.query.WriteString(",")
	}
	qb.query.WriteString(fmt.Sprintf(`{"range":{"%s":{"gte":"%s","lte":"%s"}}}`, field, lower, upper))
	qb.queryCount++
	return qb
}

func (qb *QueryBuilder) Should() *QueryBuilder {
	qb.query.WriteString(`]}},{"bool":{"must":[`)
	qb.queryCount = 0
	return qb
}

func (qb *QueryBuilder) AddSort(field string, order string) *QueryBuilder {
	if qb.sortCount > 0 {
		qb.sort.WriteString(",")
	}
	qb.sort.WriteString(fmt.Sprintf(`{"%s":{"order":"%s"}}`, field, order))
	qb.sortCount++
	return qb
}

func (qb *QueryBuilder) AddSize(size int) *QueryBuilder {
	qb.size = size
	return qb
}

func (qb *QueryBuilder) SetPage(page int) *QueryBuilder {
	qb.page = page
	return qb
}

func (qb *QueryBuilder) Query() ([]types.Hit, *errors.AppError) {
	offset := (qb.page - 1) * qb.size
	qb.sort.WriteString(`]`)
	qb.query.WriteString(fmt.Sprintf(`]}}]}},"size":"%d","from":"%d",%s}`,
		qb.size,
		offset,
		qb.sort.String(),
	))
	log.Println("[Query] Query:", qb.query.String())
	res, err := qb.client.
		Search().
		Index(qb.index).
		Raw(strings.NewReader(qb.query.String())).
		Do(context.Background())

	if err != nil {
		log.Printf("[Query] Error getting response: %s", err)
		return nil, errors.NewAppError(500, err.Error())
	}
	return res.Hits.Hits, nil
}
