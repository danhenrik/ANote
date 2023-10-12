package viewmodels

import (
	"anote/internal/domain"
	"time"
)

type CreateNoteVM struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

type NoteVM struct {
	Id        string
	Title     string
	Content   string
	AuthorID  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (note CreateNoteVM) ToDomainNote() domain.Note {
	return domain.Note{
		Title:   note.Title,
		Content: note.Content,
	}
}
