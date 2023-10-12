package interfaces

import (
	"anote/internal/domain"
	"anote/internal/errors"
)

type NoteRepository interface {
	Create(note *domain.Note) *errors.AppError

	GetByID(id string) (*domain.Note, *errors.AppError)
	GetByTitle(title string) ([]domain.Note, *errors.AppError)

	Update(user *domain.Note) *errors.AppError
	Delete(username string) *errors.AppError
}
