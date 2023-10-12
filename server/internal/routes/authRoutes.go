package routes

import (
	httpAdapter "anote/cmd/interfaces"
	"anote/internal/container"
	"anote/internal/viewmodels"
	"encoding/json"
	"log"
)

func Login(request httpAdapter.Request) httpAdapter.Response {
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

func RequestPasswordReset(request httpAdapter.Request) httpAdapter.Response {
	// generate password reset token

	// save token to database

	// send password reset token to email via link

	return httpAdapter.NewNoContentRespone()
}

func ChangePassword(request httpAdapter.Request) httpAdapter.Response {
	// check token validity

	// reset password

	return httpAdapter.NewNoContentRespone()
}
