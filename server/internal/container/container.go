package container

import (
	"anote/internal/constants"
	"anote/internal/helpers"
	"anote/internal/interfaces"
	IRepo "anote/internal/interfaces/repositories"
	"anote/internal/repositories"
	"anote/internal/services"
	"anote/internal/storage/database"
)

// This is a DI container
func Config() {
	constants.Config()

	DBConn = database.GetConnection()
	JwtProvider = helpers.NewJwtProvider()

	UserRepository = repositories.NewUserRepository(DBConn)
	NoteTagRepository = repositories.NewNoteTagRepository(DBConn)
	NoteRepository = repositories.NewNoteRepository(DBConn)

	UserService = services.NewUserService(UserRepository)
	AuthService = services.NewAuthService(UserRepository, JwtProvider)
	NoteService = services.NewNoteService(UserRepository, NoteRepository, NoteTagRepository)
	NoteTagService = services.NewNoteTagService(NoteTagRepository)
}

var DBConn interfaces.DBConnection
var JwtProvider interfaces.JwtProvider

var UserRepository IRepo.UserRepository

var NoteTagRepository IRepo.NoteTagRepository
var NoteRepository IRepo.NoteRepository

var UserService services.UserService
var AuthService services.AuthService
var NoteService services.NoteService
var NoteTagService services.NoteTagService
