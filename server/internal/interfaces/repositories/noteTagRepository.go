package interfaces

import (
	"anote/internal/domain"
	"anote/internal/errors"
)

type NoteTagRepository interface {
	Create(noteTag *domain.NoteTag) *errors.AppError
	GetAll() ([]domain.NoteTag, *errors.AppError)
	GetById(id string) (*domain.NoteTag, *errors.AppError)
	GetByNoteId(noteId string) ([]domain.NoteTag, *errors.AppError)
	Delete(id string) *errors.AppError
}
