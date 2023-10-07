package container

import (
	"anote/internal/database"
	"anote/internal/helpers"
	"anote/internal/repositories"
	"anote/internal/services"
)

// this file is used to setup the dependencies of the application

// Setup DB Connection
var DBConn = database.GetConnection()
var JwtProvider = helpers.NewJwtProvider()

// Setup Repositories
var UserRepository = repositories.NewUserRepository(DBConn)

// Setup Services
var UserService = services.NewUserService(UserRepository)
var AuthService = services.NewAuthService(UserRepository, JwtProvider)
