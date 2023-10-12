package services

import (
	"anote/internal/domain"
	"anote/internal/errors"
	"anote/internal/helpers"
	IRepo "anote/internal/interfaces/repositories"
	"anote/internal/viewmodels"
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
	hashedPassword, err := helpers.Hash(user.Password)
	if err != nil {
		return errors.NewAppError(500, "Internal server error")
	}
	user.Password = hashedPassword

	if err := this.userRepository.Create(user); err != nil {
		return err
	}
	return nil
}

func (this UserService) GetAll() ([]domain.User, *errors.AppError) {
	users, err := this.userRepository.GetAll()
	if err != nil {
		return []domain.User{}, nil
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

func (_ UserService) Update(user viewmodels.UserVM) (domain.User, *errors.AppError) {
	return domain.User{}, nil
}

func (_ UserService) Delete(username string) (bool, *errors.AppError) {
	return true, nil
}
