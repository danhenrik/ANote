package main

import (
	httpAdapter "anote/cmd/interfaces"
	"anote/internal/config"
	"anote/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/users", httpAdapter.NewGinAdapter(routes.CreateUserController))
	r.GET("/users", httpAdapter.NewGinAdapter(routes.GetAllUsersController))
	r.GET("/users/username/:username", httpAdapter.NewGinAdapter(routes.GetUserByUsernameController))
	r.GET("/users/email/:email", httpAdapter.NewGinAdapter(routes.GetUserByEmailController))
	r.POST("/login", httpAdapter.NewGinAdapter(routes.Login))

	r.Run("localhost:" + config.PORT)
}
