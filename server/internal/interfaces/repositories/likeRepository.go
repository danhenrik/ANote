package interfaces

import (
	"anote/internal/domain"
	"anote/internal/errors"
)

type LikeRepository interface {
	Create(like *domain.Like) *errors.AppError
	Delete(idUser string, idNote string) *errors.AppError
}
