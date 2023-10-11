package container

import (
	"anote/internal/constants"
	"anote/internal/database"
	"anote/internal/helpers"
	"anote/internal/ports"
	IRepo "anote/internal/ports/repositories"
	"anote/internal/repositories"
	"anote/internal/services"
)

// This is a DI container
func Config() {
	constants.Config()

	DBConn = database.GetConnection()
	JwtProvider = helpers.NewJwtProvider()

	UserRepository = repositories.NewUserRepository(DBConn)
	// NoteTagRepository = repositories.NewNoteTagRepository(DBConn)
	NoteRepository = repositories.NewNoteRepository(DBConn)

	UserService = services.NewUserService(UserRepository)
	AuthService = services.NewAuthService(UserRepository, JwtProvider)
	// NoteService = services.NewNoteTagService(NoteTagRepository, UserRepository, NoteRepository)
	// NoteService = services.NewNoteService(NoteTagRepository, UserRepository, NoteRepository)
}

var DBConn ports.DBConnection
var JwtProvider ports.JwtProvider

var UserRepository IRepo.UserRepository

// var NoteTagRepository IRepo.NoteTagRepository
var NoteRepository IRepo.NoteRepository

var UserService services.UserService
var AuthService services.AuthService

// var NoteTagService services.NoteTagService
// var NoteService services.NoteService
