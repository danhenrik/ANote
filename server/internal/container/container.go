package container

import (
	"anote/internal/constants"
	"anote/internal/database"
	"anote/internal/helpers"
	"anote/internal/ports"
	"anote/internal/repositories"
	"anote/internal/services"
)

// This is a DI container
func Config() {
	constants.Config()

	DBConn = database.GetConnection()
	JwtProvider = helpers.NewJwtProvider()

	UserRepository = repositories.NewUserRepository(DBConn)

	UserService = services.NewUserService(UserRepository)
	AuthService = services.NewAuthService(UserRepository, JwtProvider)
}

var DBConn ports.DBConnection
var JwtProvider ports.JwtProvider

var UserRepository ports.UserRepository

var UserService services.UserService
var AuthService services.AuthService
