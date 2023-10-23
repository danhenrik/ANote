package interfaces

import (
	"anote/internal/domain"
	"anote/internal/errors"
)

type AuthRepository interface {
	SaveToken(token string, userId string) *errors.AppError
	RetrieveToken(token string) (*domain.User, *errors.AppError)
	DeleteToken(token string) *errors.AppError
}
