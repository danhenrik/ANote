package services

import (
	"anote/internal/helpers"
	"anote/internal/ports"
	IRepo "anote/internal/ports/repositories"
	errors "anote/internal/types"
	"anote/internal/viewmodels"
	"log"
)

type AuthService struct {
	userRepository IRepo.UserRepository
	jwtProvider    ports.JwtProvider
}

func NewAuthService(
	userRepository IRepo.UserRepository,
	JwtProvider ports.JwtProvider,
) AuthService {
	return AuthService{
		userRepository: userRepository,
		jwtProvider:    JwtProvider,
	}
}

func (this AuthService) Login(login viewmodels.LoginVM) (string, *errors.AppError) {
	userFromDB, err := this.userRepository.GetUserWithPassword(login.Login)
	if err != nil {
		log.Println("[Login] Error on get user:", err)
		return "", err
	}
	if userFromDB == nil {
		return "", errors.NewAppError(400, "User not found")
	}

	if passwordMatch := helpers.CheckHash(login.Password, userFromDB.Password); !passwordMatch {
		return "", errors.NewAppError(400, "Invalid password")
	}

	jwt, e := this.jwtProvider.CreateToken(userFromDB)
	if e != nil {
		log.Println("[Login] Error on JWT creation:", e)
		return "", errors.NewAppError(500, "Error on JWT creation")
	}
	return jwt, nil
}
