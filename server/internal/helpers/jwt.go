package helpers

import (
	"anote/internal/constants"
	"anote/internal/domain"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtHelper struct{}

func NewJwtProvider() JwtHelper {
	return JwtHelper{}
}

func (_ JwtHelper) CreateToken(user *domain.User) (string, error) {
	claims := jwt.MapClaims{
		"id":  user.Id,
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.JWT_SECRET))
}

func (_ JwtHelper) ValidateToken(token string) (bool, string) {
	parsed, err := jwt.Parse(
		token,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(constants.JWT_SECRET), nil
		},
		jwt.WithValidMethods([]string{"HS256"}),
	)
	if err != nil || !parsed.Valid {
		return false, ""
	}

	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return false, ""
	}

	rawId, found := claims["id"]
	if !found {
		return false, ""
	}
	userId, ok := rawId.(string)
	if !ok {
		return false, ""
	}
	return true, userId
}
