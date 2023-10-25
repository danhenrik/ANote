package es

import (
	"anote/internal/errors"
	"anote/internal/interfaces"
	"fmt"
	"log"
	"strings"
)

type QueryBuilder struct {
	index      string
	query      *strings.Builder
	sort       *strings.Builder
	client     interfaces.ESClient
	queryCount int
	sortCount  int
	size       int
	page       int
}

func NewQueryBuilder(indexName string, esClient interfaces.ESClient) *QueryBuilder {
	query := strings.Builder{}
	sort := strings.Builder{}

	query.WriteString(`{"query":{"bool":{"should": [{"bool":{"must":[`)
	sort.WriteString(`"sort":[`)
	return &QueryBuilder{
		index:      indexName,
		client:     esClient,
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

func (qb *QueryBuilder) Query() ([]any, *errors.AppError) {
	offset := (qb.page - 1) * qb.size
	qb.sort.WriteString(`]`)
	qb.query.WriteString(fmt.Sprintf(`]}}]}},"size":"%d","from":"%d",%s}`,
		qb.size,
		offset,
		qb.sort.String(),
	))
	log.Println("[Query] Query:", qb.query.String())
	res, err := qb.client.Search(qb.index, qb.query.String())

	if err != nil {
		log.Printf("[Query] Error getting response: %s", err)
		return nil, errors.NewAppError(500, err.Error())
	}
	return res, nil
}
