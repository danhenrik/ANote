package services

import (
	"anote/internal/domain"
	"anote/internal/errors"
	"anote/internal/helpers"
	IRepo "anote/internal/interfaces/repositories"
	"io"
	"os"
	"strings"
)

type UserService struct {
	userRepository IRepo.UserRepository
}

func NewUserService(userRepository IRepo.UserRepository) UserService {
	return UserService{
		userRepository: userRepository,
	}
}

func (this UserService) Create(user *domain.User) *errors.AppError {
	if isValidEmail := helpers.ValidateEmail(user.Email); !isValidEmail {
		return errors.NewAppError(400, "Invalid email")
	}
	hashedPassword, err := helpers.Hash(*user.Password)
	if err != nil {
		return errors.NewAppError(500, "Internal server error")
	}
	user.Password = &hashedPassword

	if err := this.userRepository.Create(user); err != nil {
		return err
	}
	return nil
}

func (this UserService) SaveAvatar(userId string, filename string) *errors.AppError {
	splitted := strings.Split(filename, ".")
	ext := splitted[len(splitted)-1]
	avatarFilename := userId + "." + ext

	// create permanent file
	file, err := os.OpenFile("internal/assets/"+avatarFilename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return errors.NewAppError(500, "Internal server error")
	}
	defer file.Close()

	// open temporary file
	tempFile, err := os.OpenFile("internal/tmp/"+filename, os.O_RDONLY, 0777)
	if err != nil {
		return errors.NewAppError(500, "Internal server error")
	}
	defer tempFile.Close()

	// move (copy + delete)
	_, err = io.Copy(file, tempFile)
	if err != nil {
		return errors.NewAppError(500, "Internal server error")
	}
	os.Remove("internal/tmp/" + filename)

	this.userRepository.SetAvatar(userId, avatarFilename)
	return nil
}

func (this UserService) DeleteAvatar(userId string) *errors.AppError {
	user, err := this.userRepository.GetByUsername(userId)
	if err != nil {
		return err
	}

	filename := *user.Avatar
	osErr := os.Remove("internal/assets/" + filename)
	if osErr != nil {
		return errors.NewAppError(500, "Error deleting file")
	}

	err = this.userRepository.SetAvatar(userId, "")
	if err != nil {
		return err
	}
	return nil
}

func (this UserService) GetAll() ([]domain.User, *errors.AppError) {
	users, err := this.userRepository.GetAll()
	if err != nil {
		return []domain.User{}, err
	}
	return users, err
}

func (this UserService) GetByUsername(username string) (*domain.User, *errors.AppError) {
	user, err := this.userRepository.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (this UserService) GetByEmail(email string) (*domain.User, *errors.AppError) {
	user, err := this.userRepository.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (this UserService) UpdateEmail(username string, email string) *errors.AppError {
	err := this.userRepository.UpdateEmail(username, email)
	if err != nil {
		return err
	}
	return nil
}

func (this UserService) UpdatePassword(
	username string,
	oldPassword string,
	newPassword string,
) *errors.AppError {
	user, err := this.userRepository.GetUserWithPassword(username)
	if err != nil {
		return err
	}
	if isValidPassword := helpers.CheckHash(oldPassword, *user.Password); !isValidPassword {
		return errors.NewAppError(400, "Invalid password")
	}

	err = this.userRepository.UpdatePassword(username, newPassword)
	if err != nil {
		return err
	}
	return nil
}

func (this UserService) Delete(username string) *errors.AppError {
	err := this.userRepository.Delete(username)
	if err != nil {
		return err
	}
	return nil
}
