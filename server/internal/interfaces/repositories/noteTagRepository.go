package interfaces

import (
	"anote/internal/domain"
	"anote/internal/errors"
)

type NoteTagRepository interface {
	Create(noteTag *domain.NoteTag) *errors.AppError
	GetById(id string) (*domain.NoteTag, *errors.AppError)
}
