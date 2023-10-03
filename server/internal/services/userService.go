package services

import (
	"anote/internal/domain"
	interfaces "anote/internal/interfaces/repositories"
	"anote/internal/viewmodels"
	utils "anote/pkg"
	"errors"
)

type UserService struct {
	UserRepository interfaces.UserRepository
}

func NewUserService(userRepository interfaces.UserRepository) UserService {
	return UserService{UserRepository: userRepository}
}

func (this UserService) Create(userVM viewmodels.UserVM) error {
	if isValidEmail := utils.ValidateEmail(userVM.Email); !isValidEmail {
		return errors.New("Invalid email")
	}

	hashedPassword, err := utils.Hash(userVM.Password)
	if err != nil {
		return err
	}
	userVM.Password = hashedPassword

	user := userVM.ToDomainUser()
	if err = this.UserRepository.Create(user); err != nil {
		return err
	}
	return nil
}

func (this UserService) GetByUsername(username string) (viewmodels.UserVM, error) {
	user, err := this.UserRepository.GetByUsername(username)
	if err != nil {
		return viewmodels.UserVM{}, err
	}

	userVM := viewmodels.UserVMFromDomainUser(user)
	return userVM, nil
}

func (this UserService) GetByEmail(email string) (viewmodels.UserVM, error) {
	user, err := this.UserRepository.GetByEmail(email)
	if err != nil {
		return viewmodels.UserVM{}, err
	}

	userVM := viewmodels.UserVMFromDomainUser(user)
	return userVM, nil
}

func (_ UserService) GetAll() ([]domain.User, error) {
	return []domain.User{}, nil
}

func (_ UserService) Update(user domain.User) (domain.User, error) {
	return domain.User{}, nil
}

func (_ UserService) Delete(username string) (bool, error) {
	return true, nil
}
