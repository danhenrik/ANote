package main

import (
	httpAdapter "anote/cmd/interfaces"
	"anote/internal/constants"
	"anote/internal/container"
	"anote/internal/middlewares"
	"anote/internal/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	container.Config()
}

// func main() {
// 	r := gin.Default()
// 	r.MaxMultipartMemory = constants.MAX_MULTIPART_SIZE
// 	r.POST("/test", func(ctx *gin.Context) {
// 		file, _ := ctx.FormFile("file")
// 		log.Println(file.Filename)

// 		ctx.SaveUploadedFile(file, "./internal/tmp/"+file.Filename)
// 		ctx.String(201, fmt.Sprintf("'%s' uploaded!", file.Filename))
// 	})
// 	r.Run()
// }

func main() {
	r := gin.Default()
	r.MaxMultipartMemory = constants.MAX_MULTIPART_SIZE

	r.GET("/static/*filepath", func(c *gin.Context) {
		c.File("./internal/assets/" + c.Param("filepath"))
	})

	// users endpoints
	r.POST("/users", httpAdapter.NewGinAdapter(
		routes.CreateUserController,
		middlewares.ParseMultipart,
		middlewares.SaveFile("avatar", []string{"png", "jpg", "jpeg"}),
	))
	r.DELETE("/users/avatar", httpAdapter.NewGinAdapter(routes.DeleteUserAvatarController, middlewares.AuthenticateUser))
	r.PATCH("/users/avatar", httpAdapter.NewGinAdapter(
		routes.UpdateUserAvatarController,
		middlewares.AuthenticateUser,
		middlewares.ParseMultipart,
		middlewares.SaveFile("avatar", []string{"png", "jpg", "jpeg"}),
	))
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
	r.POST("/communities", httpAdapter.NewGinAdapter(
		routes.CreateCommunityController,
		middlewares.AuthenticateUser,
		middlewares.ParseMultipart,
		middlewares.SaveFile("background", []string{"png", "jpg", "jpeg"}),
	))
	r.DELETE("/communities/background/:id", httpAdapter.NewGinAdapter(routes.DeleteCommunityBackgroundController, middlewares.AuthenticateUser))
	r.PATCH("/communities/background/:id", httpAdapter.NewGinAdapter(
		routes.UpdateCommunityBackgroundController,
		middlewares.AuthenticateUser,
		middlewares.ParseMultipart,
		middlewares.SaveFile("background", []string{"png", "jpg", "jpeg"}),
	))
	r.POST("/communities/join/:id", httpAdapter.NewGinAdapter(routes.JoinCommunityController, middlewares.AuthenticateUser))
	r.POST("/communities/leave/:id", httpAdapter.NewGinAdapter(routes.LeaveCommunityController, middlewares.AuthenticateUser))
	r.GET("/communities", httpAdapter.NewGinAdapter(routes.GetAllCommunitiesController))
	r.GET("/communities/my", httpAdapter.NewGinAdapter(routes.GetCurrentUserCommunities, middlewares.AuthenticateUser))
	r.DELETE("/communities/:id", httpAdapter.NewGinAdapter(routes.DeleteCommunityController, middlewares.AuthenticateUser))

	// likes endpoints
	r.GET("/likes/:idUser/:idNote", httpAdapter.NewGinAdapter(routes.GetLikeController, middlewares.AuthenticateUser))
	r.GET("/likes/count/:idNote", httpAdapter.NewGinAdapter(routes.CountLikeByIdNoteController, middlewares.AuthenticateUser))
	r.POST("/likes", httpAdapter.NewGinAdapter(routes.CreateLikeController, middlewares.AuthenticateUser))
	r.DELETE("/likes/:idUser/:idNote", httpAdapter.NewGinAdapter(routes.DeleteLikeController, middlewares.AuthenticateUser))

	// comment endpoints
	r.GET("/comments/:idNote", httpAdapter.NewGinAdapter(routes.GetNoteCommentsController, middlewares.AuthenticateUser))
	r.GET("/comments/count/:idNote", httpAdapter.NewGinAdapter(routes.CountCommentByIdNoteController, middlewares.AuthenticateUser))
	r.POST("/comments", httpAdapter.NewGinAdapter(routes.CreateCommentController, middlewares.AuthenticateUser))
	r.DELETE("/comments/:id", httpAdapter.NewGinAdapter(routes.DeleteCommentController, middlewares.AuthenticateUser))

	r.Run()
}
