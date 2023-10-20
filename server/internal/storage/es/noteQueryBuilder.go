package es

type NoteQueryBuilder struct {
	QueryBuilder *QueryBuilder
}

func NewNoteQueryBuilder() *NoteQueryBuilder {
	return &NoteQueryBuilder{
		QueryBuilder: NewQueryBuilder("notes"),
	}
}

func (qb *NoteQueryBuilder) AddAuthorQuery(authorID string) *NoteQueryBuilder {
	qb.QueryBuilder.AddQuery("author_id", authorID)
	return qb
}

func (qb *NoteQueryBuilder) AddTitleQuery(title string) *NoteQueryBuilder {
	qb.QueryBuilder.AddQuery("title", title)
	return qb
}

func (qb *NoteQueryBuilder) AddContentQuery(content string) *NoteQueryBuilder {
	qb.QueryBuilder.AddQuery("content", content)
	return qb
}

func (qb *NoteQueryBuilder) AddTagQuery(tags []string) *NoteQueryBuilder {

	// qb.QueryBuilder.AddQuery("tags", tags)
	return qb
}
