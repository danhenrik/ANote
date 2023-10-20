package routes

import (
	httpAdapter "anote/cmd/interfaces"
	"anote/internal/container"
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

	jwt, err := container.AuthService.Login(loginVM)
	if err != nil {
		log.Println("[AuthController] Error on login:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}
	return httpAdapter.NewSuccessResponse(200, jwt)
}

func RequestPasswordResetController(request httpAdapter.Request) httpAdapter.Response {
	// generate password reset token

	// save token to database

	// send password reset token to email via link

	return httpAdapter.NewNoContentRespone()
}

func ChangeUserPasswordController(request httpAdapter.Request) httpAdapter.Response {
	var userVM viewmodels.ChangePasswordVM
	if err := json.Unmarshal([]byte(request.Body), &userVM); err != nil {
		log.Println("[UserController] Error on change password unmarshal:", err)
		return httpAdapter.NewErrorResponse(400, "Invalid content-type")
	}

	if err := container.UserService.ChangePassword(userVM.Token, userVM.NewPassword); err != nil {
		log.Println("[UserController] Error on change password:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}
	return httpAdapter.NewNoContentRespone()
}
