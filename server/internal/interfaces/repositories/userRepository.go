package interfaces

import (
	"anote/internal/domain"
	"anote/internal/errors"
)

type UserRepository interface {
	Create(user *domain.User) *errors.AppError
	GetAll() ([]domain.User, *errors.AppError)
	GetUserWithPassword(key string) (*domain.User, *errors.AppError)
	GetByUsername(username string) (*domain.User, *errors.AppError)
	GetByEmail(email string) (*domain.User, *errors.AppError)
	UpdateEmail(username, email string) *errors.AppError
	UpdatePassword(username, password string) *errors.AppError
	Delete(username string) *errors.AppError
}
