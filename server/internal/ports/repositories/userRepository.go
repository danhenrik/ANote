package ports

import "anote/internal/domain"

type UserRepository interface {
	Create(user domain.User) error
	GetByUsername(username string) (domain.User, error)
	GetByEmail(email string) (domain.User, error)
	GetAll() ([]domain.User, error)
	Update(user domain.User) error
	Delete(username string) error
}
