package main

import (
	httpAdapter "anote/cmd/interfaces"
	userControllers "anote/cmd/users"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/users", httpAdapter.NewGinAdapter(userControllers.CreateUserController))
	r.GET("/users", httpAdapter.NewGinAdapter(userControllers.GetUserController))

	r.Run()
}
