package container

import (
	"anote/internal/constants"
	"anote/internal/helpers"
	"anote/internal/interfaces"
	IRepo "anote/internal/interfaces/repositories"
	"anote/internal/repositories"
	"anote/internal/services"
	"anote/internal/storage/database"
	"anote/internal/storage/es"
)

// This is a DI container
func Config() {
	constants.Config()

	DBConn = database.GetConnection()
	JwtProvider = helpers.NewJwtProvider()
	ESClient = es.NewESClient()
	QueryBuilder := es.NewNoteQueryBuilder(ESClient)

	AuthRepository = repositories.NewAuthRepository(DBConn)
	CommunityRepository = repositories.NewCommunityRepository(DBConn)
	NoteRepository = repositories.NewNoteRepository(DBConn)
	NoteTagRepository = repositories.NewNoteTagRepository(DBConn)
	UserRepository = repositories.NewUserRepository(DBConn)
	LikeRepository = repositories.NewLikeRepository(DBConn)
	CommentRepository = repositories.NewCommentRepository(DBConn)

	AuthService = services.NewAuthService(AuthRepository, UserRepository, JwtProvider)
	CommunityService = services.NewCommunityService(CommunityRepository)
	NoteService = services.NewNoteService(UserRepository, CommunityRepository, NoteRepository, NoteTagRepository, QueryBuilder)
	NoteTagService = services.NewNoteTagService(NoteTagRepository)
	UserService = services.NewUserService(UserRepository)
	LikeService = services.NewLikeService(LikeRepository)
	CommentService = services.NewCommentService(CommentRepository, UserRepository)
}

var DBConn interfaces.DBConnection
var JwtProvider interfaces.JwtProvider
var ESClient interfaces.ESClient

var AuthRepository IRepo.AuthRepository
var CommunityRepository IRepo.CommunityRepository
var NoteRepository IRepo.NoteRepository
var NoteTagRepository IRepo.NoteTagRepository
var UserRepository IRepo.UserRepository
var LikeRepository IRepo.LikeRepository
var CommentRepository IRepo.CommentRepository

var AuthService services.AuthService
var CommunityService services.CommunityService
var NoteService services.NoteService
var NoteTagService services.NoteTagService
var UserService services.UserService
var LikeService services.LikeService
var CommentService services.CommentService
