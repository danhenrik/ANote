package middlewares

import (
	httpAdapter "anote/cmd/interfaces"
	"anote/internal/container"
	"anote/internal/errors"
	"log"
	"strings"
)

func AuthenticateUser(req httpAdapter.Request) (httpAdapter.Request, *errors.AppError) {
	bearerToken, found := req.GetHeader("Authorization")
	if !found {
		log.Println("[AuthMiddleware] Auth Header not found")
		return req, errors.NewAppError(401, "Couldn't locate a valid Authorization Header")
	}

	jwtToken := strings.Split(bearerToken, " ")[1]
	tokenIsValid, userId := container.JwtProvider.ValidateToken(jwtToken)
	if !tokenIsValid {
		log.Println("[AuthMiddleware] Invalid JWT Token")
		return req, errors.NewAppError(400, "Invalid JWT Token")
	}

	user, err := container.UserRepository.GetByUsername(userId)
	if err != nil {
		log.Println("[AuthMiddleware] Error on get token owner", err)
		return req, err
	}

	req.User = httpAdapter.UserIdentity{
		ID:    user.Id,
		Email: user.Email,
	}
	log.Println("[AuthMiddleware] Successfully authenticated user", user.Id)
	return req, nil
}
