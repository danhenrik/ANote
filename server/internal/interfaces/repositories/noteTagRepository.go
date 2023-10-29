package interfaces

import (
	"anote/internal/domain"
	"anote/internal/errors"
)

type NoteTagRepository interface {
	Create(noteTag *domain.Tag) *errors.AppError
	GetAll() ([]domain.Tag, *errors.AppError)
	GetById(id string) (*domain.Tag, *errors.AppError)
	GetByName(tagName string) (*domain.Tag, *errors.AppError)
	GetByNoteId(noteId string) ([]domain.Tag, *errors.AppError)
	Delete(id string) *errors.AppError
}
