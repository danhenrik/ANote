package services

import (
	"anote/internal/domain"
	"anote/internal/errors"
	"anote/internal/helpers"
	"anote/internal/interfaces"
	IRepo "anote/internal/interfaces/repositories"
	"anote/internal/viewmodels"
	"fmt"
	"log"
	"time"
)

type AuthService struct {
	authRepository IRepo.AuthRepository
	userRepository IRepo.UserRepository
	jwtProvider    interfaces.JwtProvider
}

func NewAuthService(
	authRepository IRepo.AuthRepository,
	userRepository IRepo.UserRepository,
	JwtProvider interfaces.JwtProvider,
) AuthService {
	return AuthService{
		authRepository: authRepository,
		userRepository: userRepository,
		jwtProvider:    JwtProvider,
	}
}

func (this AuthService) Login(login viewmodels.LoginVM) (string, *domain.User, *errors.AppError) {
	userFromDB, err := this.userRepository.GetUserWithPassword(login.Login)
	if err != nil {
		log.Println("[Login] Error on get user:", err)
		return "", nil, err
	}
	if userFromDB == nil {
		return "", nil, errors.NewAppError(400, "User not found")
	}

	if passwordMatch := helpers.CheckHash(login.Password, *userFromDB.Password); !passwordMatch {
		return "", nil, errors.NewAppError(400, "Invalid password")
	}

	jwt, e := this.jwtProvider.CreateToken(userFromDB)
	if e != nil {
		log.Println("[Login] Error on JWT creation:", e)
		return "", nil, errors.NewAppError(500, "Error on JWT creation")
	}
	return jwt, userFromDB, nil
}

func (this AuthService) RequestPasswordReset(email string) (string, *errors.AppError) {
	userFromDB, err := this.userRepository.GetByEmail(email)
	if err != nil {
		log.Println("[RequestPasswordReset] Error on get user:", err)
		return "", err
	}
	if userFromDB == nil {
		return "", errors.NewAppError(400, "User not found")
	}

	tokenKey := fmt.Sprint(userFromDB.Id, time.Now().UnixMilli())
	token, hashErr := helpers.Hash(tokenKey)
	if hashErr != nil {
		log.Println("[RequestPasswordReset] Error on token creation (hash):", err)
		return "", errors.NewAppError(500, "Error on token creation")
	}

	err = this.authRepository.SaveToken(token, userFromDB.Id)
	if err != nil {
		log.Println("[RequestPasswordReset] Error on save token:", err)
		return "", err
	}
	return token, nil
}

func (this AuthService) ChangePassword(token string, newPassword string) *errors.AppError {
	userFromToken, err := this.authRepository.RetrieveToken(token)
	if err != nil {
		log.Println("[ChangePassword] Error on get user by token:", err)
		return err
	}
	if userFromToken == nil {
		log.Println("[ChangePassword] Token not found")
		return errors.NewAppError(400, "Token not found")
	}

	password, hashErr := helpers.Hash(newPassword)
	if hashErr != nil {
		log.Println("[ChangePassword] Error on password hash:", err)
		return errors.NewAppError(500, "Error on password hash")
	}

	err = this.userRepository.UpdatePassword(userFromToken.Id, password)
	if err != nil {
		log.Println("[ChangePassword] Error on update password:", err)
		return err
	}

	// If it fails should explore anyways so there's no reason for returning an error
	this.authRepository.DeleteToken(token)
	return nil
}
