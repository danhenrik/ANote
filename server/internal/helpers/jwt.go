package helpers

import (
	"anote/internal/constants"
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

	return token.SignedString([]byte(constants.JWT_SECRET))
}

func (_ JwtHelper) ValidateToken(token string) (bool, string) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.JWT_SECRET), nil
	})
	if err != nil {
		return false, ""
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return false, ""
	}

	userId, found := claims["id"]
	if !found {
		return false, ""
	}
	return true, userId.(string)
}
