package services

import (
	"anote/internal/domain"
	"anote/internal/ports"
	"anote/internal/types"
	errors "anote/internal/types"
	"anote/internal/utils"
	"anote/internal/viewmodels"
)

type UserService struct {
	UserRepository ports.UserRepository
}

func NewUserService(userRepository ports.UserRepository) UserService {
	return UserService{UserRepository: userRepository}
}

func (this UserService) Create(user domain.User) *types.AppError {
	if isValidEmail := utils.ValidateEmail(user.Email); !isValidEmail {
		return errors.NewAppError(400, "Invalid email")
	}
	hashedPassword, err := utils.Hash(user.Password)
	if err != nil {
		return errors.NewAppError(500, "Internal server error")
	}
	user.Password = hashedPassword

	if err := this.UserRepository.Create(user); err != nil {
		return err
	}
	return nil
}

func (this UserService) GetAll() ([]domain.User, *types.AppError) {
	users, err := this.UserRepository.GetAll()
	if err != nil {
		return []domain.User{}, nil
	}
	return users, err
}

func (this UserService) GetByUsername(username string) (*domain.User, *types.AppError) {
	user, err := this.UserRepository.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (this UserService) GetByEmail(email string) (*domain.User, *types.AppError) {
	user, err := this.UserRepository.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (_ UserService) Update(user viewmodels.UserVM) (domain.User, *types.AppError) {
	return domain.User{}, nil
}

func (_ UserService) Delete(username string) (bool, *types.AppError) {
	return true, nil
}
