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

	AuthRepository = repositories.NewAuthRepository(DBConn)
	CommunityRepository = repositories.NewCommunityRepository(DBConn)
	NoteRepository = repositories.NewNoteRepository(DBConn)
	NoteTagRepository = repositories.NewNoteTagRepository(DBConn)
	UserRepository = repositories.NewUserRepository(DBConn)

	AuthService = services.NewAuthService(AuthRepository, UserRepository, JwtProvider)
	CommunityService = services.NewCommunityService(CommunityRepository)
	NoteService = services.NewNoteService(UserRepository, CommunityRepository, NoteRepository, NoteTagRepository)
	NoteTagService = services.NewNoteTagService(NoteTagRepository)
	UserService = services.NewUserService(UserRepository)
}

var DBConn interfaces.DBConnection
var JwtProvider interfaces.JwtProvider

var AuthRepository IRepo.AuthRepository
var CommunityRepository IRepo.CommunityRepository
var NoteRepository IRepo.NoteRepository
var NoteTagRepository IRepo.NoteTagRepository
var UserRepository IRepo.UserRepository

var AuthService services.AuthService
var CommunityService services.CommunityService
var NoteService services.NoteService
var NoteTagService services.NoteTagService
var UserService services.UserService
