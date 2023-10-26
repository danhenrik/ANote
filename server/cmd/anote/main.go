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
	r.GET("/notes/feed", httpAdapter.NewGinAdapter(routes.GetNoteFeedController, middlewares.AuthenticateUser))
	r.GET("/notes", httpAdapter.NewGinAdapter(routes.SearchNoteController))
	r.GET("/notes/:id", httpAdapter.NewGinAdapter(routes.GetNoteByIDController))
	r.GET("/notes/community/:id", httpAdapter.NewGinAdapter(routes.GetNoteByCommunityIDController))
	r.GET("/notes/author/:id", httpAdapter.NewGinAdapter(routes.GetNoteByAuthorIDController))
	r.PATCH("/notes", httpAdapter.NewGinAdapter(routes.UpdateNoteController, middlewares.AuthenticateUser))
	r.DELETE("/notes/:id", httpAdapter.NewGinAdapter(routes.DeleteNoteController, middlewares.AuthenticateUser))

	// tags endpoints
	r.POST("/tags", httpAdapter.NewGinAdapter(routes.CreateTagController, middlewares.AuthenticateUser))
	r.GET("/tags", httpAdapter.NewGinAdapter(routes.GetAllTagsController))
	r.DELETE("/tags/:id", httpAdapter.NewGinAdapter(routes.DeleteTagController, middlewares.AuthenticateUser))

	// communities endpoints
	r.POST("/communities", httpAdapter.NewGinAdapter(routes.CreateCommunityController, middlewares.AuthenticateUser))
	r.POST("/communities/join/:id", httpAdapter.NewGinAdapter(routes.JoinCommunityController, middlewares.AuthenticateUser))
	r.POST("/communities/leave/:id", httpAdapter.NewGinAdapter(routes.LeaveCommunityController, middlewares.AuthenticateUser))
	r.GET("/communities", httpAdapter.NewGinAdapter(routes.GetAllCommunitiesController))
	r.DELETE("/communities/:id", httpAdapter.NewGinAdapter(routes.DeleteCommunityController, middlewares.AuthenticateUser))

	// likes endpoints
	r.GET("/likes/:idUser/:idNote", httpAdapter.NewGinAdapter(routes.GetLikeController))
	r.GET("/likes/count/:idNote", httpAdapter.NewGinAdapter(routes.CountLikeByIdNoteController))
	r.POST("/likes", httpAdapter.NewGinAdapter(routes.CreateLikeController))
	r.DELETE("/likes/:idUser/:idNote", httpAdapter.NewGinAdapter(routes.DeleteLikeController))

	// comment endpoints
	r.GET("/comments/:idNote", httpAdapter.NewGinAdapter(routes.GetNoteCommentsController))
	r.GET("/comments/count/:idNote", httpAdapter.NewGinAdapter(routes.CountCommentByIdNoteController))
	r.POST("/comments", httpAdapter.NewGinAdapter(routes.CreateCommentController))
	r.DELETE("/comments/:id", httpAdapter.NewGinAdapter(routes.DeleteCommentController))

	r.Run()
}
