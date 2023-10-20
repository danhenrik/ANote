package domain

type Note struct {
	Id        string
	Title     string
	AuthorID  string
	Content   string
	CreatedAt string
	UpdatedAt string
}

type FullNote struct {
	Id        string
	Title     string
	AuthorID  string
	Content   string
	CreatedAt string
	UpdatedAt string
	Tags      []NoteTag
	// Communities []domain.Community
}
