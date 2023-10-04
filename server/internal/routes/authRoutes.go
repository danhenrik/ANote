package routes

import httpAdapter "anote/cmd/interfaces"

func Login(request httpAdapter.Request) httpAdapter.Response {
	response := httpAdapter.NewErrorResponse(400, "Login failed")
	return response
}
