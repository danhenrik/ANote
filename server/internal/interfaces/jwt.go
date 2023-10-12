package interfaces

import "anote/internal/domain"

type JwtProvider interface {
	CreateToken(user *domain.User) (string, error)
	ValidateToken(token string) (bool, string)
}
