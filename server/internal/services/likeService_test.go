package services

import (
	"anote/internal/domain"
	"anote/internal/errors"
	mock_interfaces "anote/internal/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// TestLikeService_Create tests the Create method of the LikeService
func TestLikeService_Create(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	likeRepository := mock_interfaces.NewMockLikeRepository(ctrl)
	likeService := NewLikeService(likeRepository)

	like := &domain.Like{
		Id:        "some_id",
		UserId:    "user_id",
		NoteId:    "note_id",
		CreatedAt: "created_at",
	}

	// Mock repository response
	likeRepository.EXPECT().GetByIdUserAndIdNote("user_id", "note_id").Return(nil, nil)
	likeRepository.EXPECT().Create(like).Return(nil)

	// Act
	err := likeService.Create(like)

	// Assert
	assert.Nil(t, err)
}

// TestLikeService_Create_FailWhenIdNotExists tests the Create method of the LikeService when the ids (UserId or NoteId) don't exist
func TestLikeService_Create_FailWhenIdNotExists(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	likeRepository := mock_interfaces.NewMockLikeRepository(ctrl)
	likeService := NewLikeService(likeRepository)

	// Mock repository response
	likeRepository.EXPECT().GetByIdUserAndIdNote("user_id", "note_id").Return(nil, errors.NewAppError(500, "Internal server error"))

	// Act
	err := likeService.Create(&domain.Like{
		Id:        "some_id",
		UserId:    "user_id",
		NoteId:    "note_id",
		CreatedAt: "created_at",
	})

	// Assert
	assert.Error(t, err)
}

// TestLikeService_Create_FailWhenLikeAlreadyExists tests the Create method of the LikeService when a like is already created for that note
func TestLikeService_Create_FailWhenLikeAlreadyExists(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	likeRepository := mock_interfaces.NewMockLikeRepository(ctrl)
	likeService := NewLikeService(likeRepository)

	like := &domain.Like{
		Id:        "some_id",
		UserId:    "user_id",
		NoteId:    "note_id",
		CreatedAt: "created_at",
	}

	// Mock repository response
	likeRepository.EXPECT().GetByIdUserAndIdNote("user_id", "note_id").Return(like, nil)

	// Act
	err := likeService.Create(like)

	// Assert
	assert.Error(t, err)
}

// TestLikeService_Create_Fail tests the Create method of the LikeService when repository fail
func TestLikeService_Create_Fail(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	likeRepository := mock_interfaces.NewMockLikeRepository(ctrl)
	likeService := NewLikeService(likeRepository)

	like := &domain.Like{
		Id:        "some_id",
		UserId:    "user_id",
		NoteId:    "note_id",
		CreatedAt: "created_at",
	}

	// Mock repository response
	likeRepository.EXPECT().GetByIdUserAndIdNote("user_id", "note_id").Return(nil, nil)
	likeRepository.EXPECT().Create(like).Return(errors.NewAppError(500, "Internal server error"))

	// Act
	err := likeService.Create(like)

	// Assert
	assert.Error(t, err)
}

// TestLikeService_Delete tests the Delete method of the LikeService
func TestLikeService_Delete(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	likeRepository := mock_interfaces.NewMockLikeRepository(ctrl)
	likeService := NewLikeService(likeRepository)

	// Mock repository response
	likeRepository.EXPECT().Delete("user_id", "note_id").Return(nil)

	// Act
	err := likeService.Delete("user_id", "note_id")

	// Assert
	assert.Nil(t, err)
}

// TestLikeService_Delete tests the Delete method of the LikeService when repository fail
func TestLikeService_Delete_Fail(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	likeRepository := mock_interfaces.NewMockLikeRepository(ctrl)
	likeService := NewLikeService(likeRepository)

	// Mock repository response
	likeRepository.EXPECT().Delete("user_id", "note_id").Return(errors.NewAppError(500, "Internal server error"))

	// Act
	err := likeService.Delete("user_id", "note_id")

	// Assert
	assert.Error(t, err)
}

// TestLikeService_GetByIdUserAndIdNote tests the GetByIdUserAndIdNote method of the LikeService
func TestLikeService_GetByIdUserAndIdNote(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	likeRepository := mock_interfaces.NewMockLikeRepository(ctrl)
	likeService := NewLikeService(likeRepository)

	like := &domain.Like{
		Id:        "some_id",
		UserId:    "user_id",
		NoteId:    "note_id",
		CreatedAt: "created_at",
	}

	// Mock repository response
	likeRepository.EXPECT().GetByIdUserAndIdNote("user_id", "note_id").Return(like, nil)

	// Act
	returnedLike, err := likeService.GetByIdUserAndIdNote("user_id", "note_id")

	// Assert
	assert.Equal(t, returnedLike, like)
	assert.Nil(t, err)
}

// TestLikeService_GetByIdUserAndIdNote_FailWhenLikeNotFound tests the GetByIdUserAndIdNote method of the LikeService when the like is not found given the ids
func TestLikeService_GetByIdUserAndIdNote_FailWhenLikeNotFound(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	likeRepository := mock_interfaces.NewMockLikeRepository(ctrl)
	likeService := NewLikeService(likeRepository)

	// Mock repository response
	likeRepository.EXPECT().GetByIdUserAndIdNote("user_id", "note_id").Return(nil, errors.NewAppError(500, "Internal server error"))

	// Act
	returnedLike, err := likeService.GetByIdUserAndIdNote("user_id", "note_id")

	// Assert
	assert.Nil(t, returnedLike)
	assert.Error(t, err)
}

// TestLikeService_CountLikeByIdNoteController tests the CountLikeByIdNoteController method of the LikeService
func TestLikeService_CountLikeByIdNoteController(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	likeRepository := mock_interfaces.NewMockLikeRepository(ctrl)
	likeService := NewLikeService(likeRepository)

	like1 := domain.Like{
		Id:        "some_id1",
		UserId:    "user_id1",
		NoteId:    "note_id1",
		CreatedAt: "created_at1",
	}

	like2 := domain.Like{
		Id:        "some_id2",
		UserId:    "user_id2",
		NoteId:    "note_id2",
		CreatedAt: "created_at2",
	}

	// Mock repository response
	likeRepository.EXPECT().GetByIdNote("like_id").Return([]domain.Like{like1, like2}, nil)

	// Act
	numberLikes, err := likeService.CountLikeByIdNoteController("like_id")

	// Assert
	assert.Equal(t, numberLikes, 2)
	assert.Nil(t, err)
}

// TestLikeService_CountLikeByIdNoteController tests the CountLikeByIdNoteController method of the LikeService when the note doesn't exist
func TestLikeService_CountLikeByIdNoteController_FailWhenNoteNotExists(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	likeRepository := mock_interfaces.NewMockLikeRepository(ctrl)
	likeService := NewLikeService(likeRepository)

	// Mock repository response
	likeRepository.EXPECT().GetByIdNote("like_id").Return(nil, errors.NewAppError(500, "Internal server error"))

	// Act
	numberLikes, err := likeService.CountLikeByIdNoteController("like_id")

	// Assert
	assert.Equal(t, numberLikes, 0)
	assert.Error(t, err)
}