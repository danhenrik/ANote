package helpers

import (
	"anote/internal/config"
	"anote/internal/domain"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtHelper struct{}

func NewJwtProvider() JwtHelper {
	return JwtHelper{}
}

func (_ JwtHelper) CreateToken(user *domain.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["id"] = user.Id

	return token.SignedString([]byte(config.JWT_SECRET))
}

func (_ JwtHelper) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_SECRET), nil
	})
	if err != nil {
		return false
	}
	return true
}
