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
	var user viewmodels.UserVM

	if err := json.Unmarshal([]byte(request.Body), &user); err != nil {
		return httpAdapter.NewErrorResponse(400, "Invalid content-type")
	}

	if err := container.UserService.Create(user); err != nil {
		log.Println("Error on create user service", err)
		return httpAdapter.NewErrorResponse(400, "Error on create user service")
	}
	return httpAdapter.NewSuccessResponse(204, nil)
}

func GetUserController(request httpAdapter.Request) httpAdapter.Response {
	userId := request.Query["username"][0]
	if userId != "" {
		user, err := container.UserService.GetByUsername(userId)
		if err != nil {
			log.Println("Error on get user by username", err)
			return httpAdapter.NewErrorResponse(http.StatusInternalServerError, err.Error())
		}
		return httpAdapter.NewSuccessResponse(http.StatusOK, user)
	}

	email := request.Query["email"][0]
	if email != "" {
		user, err := container.UserService.GetByEmail(email)
		if err != nil {
			log.Println("Error on get user by username", err)
			return httpAdapter.NewErrorResponse(http.StatusInternalServerError, err.Error())
		}
		return httpAdapter.NewSuccessResponse(http.StatusOK, user)
	}

	users, err := container.UserService.GetAll()
	if err != nil {
		log.Println("Error on get all users", err)
		return httpAdapter.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	return httpAdapter.NewSuccessResponse(http.StatusOK, users)
}
