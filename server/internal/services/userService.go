package services

import (
	ports "anote/internal/ports/repositories"
	"anote/internal/utils"
	"anote/internal/viewmodels"
	"errors"
)

type UserService struct {
	UserRepository ports.UserRepository
}

func NewUserService(userRepository ports.UserRepository) UserService {
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

func (_ UserService) GetAll() ([]viewmodels.UserVM, error) {
	return []viewmodels.UserVM{}, nil
}

func (_ UserService) Update(user viewmodels.UserVM) (viewmodels.UserVM, error) {
	return viewmodels.UserVM{}, nil
}

func (_ UserService) Delete(username string) (bool, error) {
	return true, nil
}
