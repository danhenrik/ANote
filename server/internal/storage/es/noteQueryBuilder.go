package es

import (
	"anote/internal/domain"
	"anote/internal/errors"
	"encoding/json"
	"fmt"
)

type NoteQueryBuilder struct {
	QueryBuilder *QueryBuilder
}

func NewNoteQueryBuilder() *NoteQueryBuilder {
	return &NoteQueryBuilder{
		QueryBuilder: NewQueryBuilder("notes"),
	}
}

func (qb *NoteQueryBuilder) AddAuthorQuery(authorID string) *NoteQueryBuilder {
	qb.QueryBuilder.AddMatchQuery("author.keyword", authorID)
	return qb
}

func (qb *NoteQueryBuilder) AddTitleQuery(title string) *NoteQueryBuilder {
	wildcard := fmt.Sprintf("*%s*", title)
	qb.QueryBuilder.AddWildcardQuery("title.keyword", wildcard)
	return qb
}

func (qb *NoteQueryBuilder) AddContentQuery(content string) *NoteQueryBuilder {
	wildcard := fmt.Sprintf("*%s*", content)
	qb.QueryBuilder.AddWildcardQuery("content.keyword", wildcard)
	return qb
}

func (qb *NoteQueryBuilder) AddTagsQuery(tags []string) *NoteQueryBuilder {
	qb.QueryBuilder.AddIncludeQuery("tags.name.keyword", tags)
	return qb
}

func (qb *NoteQueryBuilder) AddCommunitiesQuery(communities []string) *NoteQueryBuilder {
	qb.QueryBuilder.AddIncludeQuery("communities.name.keyword", communities)
	return qb
}

func (qb *NoteQueryBuilder) AddCreatedAtMatchQuery(date string) *NoteQueryBuilder {
	qb.QueryBuilder.AddMatchQuery("published_date", date)
	return qb
}

func (qb *NoteQueryBuilder) AddCreatedAtRangeQuery(lowerBound string, upperBound string) *NoteQueryBuilder {
	qb.QueryBuilder.AddRangeQuery("published_date", lowerBound, upperBound)
	return qb
}

func (qb *NoteQueryBuilder) AddLikesQuery(likes []string) *NoteQueryBuilder {
	qb.QueryBuilder.AddIncludeQuery("likes.user_id.keyword", likes)
	return qb
}

func (qb *NoteQueryBuilder) AddCommentersQuery(commenters []string) *NoteQueryBuilder {
	qb.QueryBuilder.AddIncludeQuery("comments.user_id.keyword", commenters)
	return qb
}

func (qb *NoteQueryBuilder) Query() ([]domain.FilteredNote, *errors.AppError) {
	noteArr, err := qb.QueryBuilder.Query()
	if err != nil {
		return nil, err
	}
	if len(noteArr) == 0 {
		return nil, nil
	}
	notes := []domain.FilteredNote{}
	for _, n := range noteArr {
		note := domain.FilteredNote{}
		err := json.Unmarshal(n.Source_, &note)
		if err != nil {
			return nil, errors.NewAppError(500, "Could not parse query result")
		}
		notes = append(notes, note)
	}
	return notes, nil
}
