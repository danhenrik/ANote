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

	// users endpoints
	r.POST("/users", httpAdapter.NewGinAdapter(routes.CreateUserController))
	r.GET("/users", httpAdapter.NewGinAdapter(routes.GetAllUsersController))
	r.GET("/users/username/:username", httpAdapter.NewGinAdapter(routes.GetUserByUsernameController))
	r.GET("/users/email/:email", httpAdapter.NewGinAdapter(routes.GetUserByEmailController))
	r.PATCH("/users/email", httpAdapter.NewGinAdapter(routes.UpdateUserEmailController, middlewares.AuthenticateUser))
	r.PATCH("/users/password", httpAdapter.NewGinAdapter(routes.UpdateUserPasswordController, middlewares.AuthenticateUser))
	r.DELETE("/users", httpAdapter.NewGinAdapter(routes.DeleteUserController, middlewares.AuthenticateUser))

	// auth endpoints
	r.POST("/auth/login", httpAdapter.NewGinAdapter(routes.LoginController))
	r.POST("/auth/request-password-reset", httpAdapter.NewGinAdapter(routes.RequestPasswordResetController))
	r.POST("/auth/change-password", httpAdapter.NewGinAdapter(routes.ChangeUserPasswordController))

	// notes endpoints
	r.POST("/notes", httpAdapter.NewGinAdapter(routes.CreateNoteController, middlewares.AuthenticateUser))
	r.GET("/notes/:id", httpAdapter.NewGinAdapter(routes.GetNoteByIDController))
	r.GET("/notes", httpAdapter.NewGinAdapter(routes.SearchNoteController))
	r.PATCH("/notes", httpAdapter.NewGinAdapter(routes.UpdateNoteController, middlewares.AuthenticateUser))
	r.DELETE("/notes/:id", httpAdapter.NewGinAdapter(routes.DeleteNoteController, middlewares.AuthenticateUser))

	// tags endpoints
	r.POST("/tags", httpAdapter.NewGinAdapter(routes.CreateTagController, middlewares.AuthenticateUser))
	r.GET("/tags", httpAdapter.NewGinAdapter(routes.GetTagsController))
	r.DELETE("/tags/:id", httpAdapter.NewGinAdapter(routes.DeleteTagController))

	r.Run()
}
