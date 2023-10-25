package interfaces

import (
	"anote/internal/domain"
	"anote/internal/errors"
)

type CommentRepository interface {
	Create(comment *domain.Comment) *errors.AppError
	Delete(idComment string) *errors.AppError
	GetNoteComments(idNote string) ([]domain.Comment, *errors.AppError)
}
