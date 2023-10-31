package routes

import (
	httpAdapter "anote/cmd/interfaces"
	"anote/internal/container"
	"anote/internal/domain"
	"anote/internal/viewmodels"
	"encoding/json"
	"log"
)

func LoginController(request httpAdapter.Request) httpAdapter.Response {
	var loginVM viewmodels.LoginVM
	if err := json.Unmarshal([]byte(request.Body), &loginVM); err != nil {
		log.Println("[AuthController] Error on login unmarshal:", err)
		return httpAdapter.NewErrorResponse(400, "Invalid content-type")
	}

	jwt, user, err := container.AuthService.Login(loginVM)
	if err != nil {
		log.Println("[AuthController] Error on login:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}

	type response struct {
		Jwt  string
		User *domain.User
	}

	return httpAdapter.NewSuccessResponse(200, response{jwt, user})
}

func RequestPasswordResetController(request httpAdapter.Request) httpAdapter.Response {
	var userVM viewmodels.RequestPasswordResetVM
	if err := json.Unmarshal([]byte(request.Body), &userVM); err != nil {
		log.Println("[UserController] Error on request password reset unmarshal:", err)
		return httpAdapter.NewErrorResponse(400, "Invalid content-type")
	}

	token, err := container.AuthService.RequestPasswordReset(userVM.Email)
	if err != nil {
		log.Println("[UserController] Error on request password reset:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}
	return httpAdapter.NewSuccessResponse(200, token)
}

func ChangeUserPasswordController(request httpAdapter.Request) httpAdapter.Response {
	var userVM viewmodels.ChangePasswordVM
	if err := json.Unmarshal([]byte(request.Body), &userVM); err != nil {
		log.Println("[UserController] Error on change password unmarshal:", err)
		return httpAdapter.NewErrorResponse(400, "Invalid content-type")
	}

	if err := container.AuthService.ChangePassword(userVM.Token, userVM.NewPassword); err != nil {
		log.Println("[UserController] Error on change password:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}
	return httpAdapter.NewNoContentResponse()
}
