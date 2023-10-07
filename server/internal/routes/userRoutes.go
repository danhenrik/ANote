package routes

import (
	httpAdapter "anote/cmd/interfaces"
	"anote/internal/container"
	"anote/internal/viewmodels"
	"encoding/json"
	"log"
	"net/http"
)

func CreateUserController(request httpAdapter.Request) httpAdapter.Response {
	var userVM viewmodels.CreateUserVM
	if err := json.Unmarshal([]byte(request.Body), &userVM); err != nil {
		log.Println("[UserController] Error on create user unmarshal:", err)
		return httpAdapter.NewErrorResponse(400, "Invalid content-type")
	}

	user := userVM.ToDomainUser()
	if err := container.UserService.Create(user); err != nil {
		log.Println("[UserController] Error on create user:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}
	return httpAdapter.NewSuccessResponse(204, nil)
}

func GetAllUsersController(request httpAdapter.Request) httpAdapter.Response {
	users, err := container.UserService.GetAll()
	if err != nil {
		log.Println("[UserController] Error on get all users:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}

	var userVMs []viewmodels.UserVM
	for _, user := range users {
		userVMs = append(userVMs, viewmodels.UserVMFromDomainUser(user))
	}
	return httpAdapter.NewSuccessResponse(http.StatusOK, userVMs)
}

func GetUserByUsernameController(request httpAdapter.Request) httpAdapter.Response {
	var userVM viewmodels.UserVM
	userId, ok := request.GetSingleParam("username")
	if ok && userId != "" {
		user, err := container.UserService.GetByUsername(userId)
		if err != nil {
			log.Println("[UserController] Error on get user by username:", err)
			return httpAdapter.NewErrorResponse(err.Status, err.Message)
		}
		userVM = viewmodels.UserVMFromDomainUser(*user)
	}
	return httpAdapter.NewSuccessResponse(http.StatusOK, userVM)
}

func GetUserByEmailController(request httpAdapter.Request) httpAdapter.Response {
	var userVM viewmodels.UserVM
	email, ok := request.GetSingleParam("email")
	if ok && email != "" {
		user, err := container.UserService.GetByEmail(email)
		if err != nil {
			log.Println("[UserController] Error on get user by email:", err)
			return httpAdapter.NewErrorResponse(err.Status, err.Message)
		}
		userVM = viewmodels.UserVMFromDomainUser(*user)
	}
	return httpAdapter.NewSuccessResponse(http.StatusOK, userVM)
}
