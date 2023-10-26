package interfaces

import (
	"anote/internal/domain"
	"anote/internal/errors"
)

type LikeRepository interface {
	Create(like *domain.Like) *errors.AppError
	Delete(idUser string, idNote string) *errors.AppError
	GetByIdUserAndIdNote(idUser string, idNote string) (*domain.Like, *errors.AppError)
}
