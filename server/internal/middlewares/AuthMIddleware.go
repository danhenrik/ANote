package middlewares

import (
	httpAdapter "anote/cmd/interfaces"
	"anote/internal/container"
	errors "anote/internal/types"
	"log"
	"strings"
)

// TODO: Add logs
func AuthenticateUser(req httpAdapter.Request) (httpAdapter.Request, *errors.AppError) {
	bearerToken, found := req.GetHeader("Authorization")
	if !found {
		return req, errors.NewAppError(401, "Couldn't locate a valid Authorization Header")
	}

	jwtToken := strings.Split(bearerToken, " ")[1]
	tokenIsValid, userId := container.JwtProvider.ValidateToken(jwtToken)
	if !tokenIsValid {
		return req, errors.NewAppError(400, "Invalid JWT Token")
	}

	log.Println(userId)
	user, err := container.UserRepository.GetByUsername(userId)
	if err != nil {
		log.Println("Error on get token owner", err)
		return req, err
	}

	req.User = httpAdapter.UserIdentity{
		ID:    user.Id,
		Email: user.Email,
	}
	return req, nil
}
