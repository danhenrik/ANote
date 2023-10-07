package ports

import (
	"anote/internal/domain"
	errors "anote/internal/types"
)

type UserRepository interface {
	Create(user domain.User) *errors.AppError
	GetUserWithPassword(key string) (*domain.User, *errors.AppError)
	GetByUsername(username string) (*domain.User, *errors.AppError)
	GetByEmail(email string) (*domain.User, *errors.AppError)
	GetAll() ([]domain.User, *errors.AppError)
	Update(user domain.User) *errors.AppError
	Delete(username string) *errors.AppError
}
