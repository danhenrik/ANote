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
