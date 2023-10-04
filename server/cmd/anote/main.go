package main

import (
	httpAdapter "anote/cmd/interfaces"
	"anote/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/users", httpAdapter.NewGinAdapter(routes.CreateUserController))
	r.GET("/users", httpAdapter.NewGinAdapter(routes.GetUserController))
	r.POST("/login", httpAdapter.NewGinAdapter(routes.Login))

	r.Run()
}
