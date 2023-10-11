package main

import (
	httpAdapter "anote/cmd/interfaces"
	"anote/internal/container"
	"anote/internal/middlewares"
	"anote/internal/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	container.Config()
}

func main() {
	r := gin.Default()

	r.POST("/users", httpAdapter.NewGinAdapter(routes.CreateUserController))
	r.GET("/users", httpAdapter.NewGinAdapter(routes.GetAllUsersController, middlewares.AuthenticateUser))
	r.GET("/users/username/:username", httpAdapter.NewGinAdapter(routes.GetUserByUsernameController))
	r.GET("/users/email/:email", httpAdapter.NewGinAdapter(routes.GetUserByEmailController))
	r.POST("/auth/login", httpAdapter.NewGinAdapter(routes.Login))

	r.Run()
}
