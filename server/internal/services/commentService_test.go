package services

import (
	"anote/internal/domain"
	"anote/internal/errors"
	mock_interfaces "anote/internal/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// TestCommentService_Create tests the Create method of the CommentService
func TestCommentService_Create(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	commentRepository := mock_interfaces.NewMockCommentRepository(ctrl)
	userRepository := mock_interfaces.NewMockUserRepository(ctrl)
	commentService := NewCommentService(commentRepository, userRepository)

	comment := &domain.Comment{
		Id:       "some_id",
		UserId:   "user_id",
		NoteId:    "note_id",
		Content:   "content",
		CreatedAt: "created_at",
	}
	// Mock repository response
	commentRepository.EXPECT().Create(comment).Return(nil)

	// Act
	err := commentService.Create(comment)

	// Assert
	assert.Nil(t, err)
}

// TestCommentService_Create_Fail tests the Create method of the CommentService when repository fail
func TestCommentService_Create_Fail(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	commentRepository := mock_interfaces.NewMockCommentRepository(ctrl)
	userRepository := mock_interfaces.NewMockUserRepository(ctrl)
	commentService := NewCommentService(commentRepository, userRepository)

	comment := &domain.Comment{
		Id:       "some_id",
		UserId:   "user_id",
		NoteId:    "note_id",
		Content:   "content",
		CreatedAt: "created_at",
	}
	// Mock repository response
	commentRepository.EXPECT().Create(comment).Return(errors.NewAppError(500, "Internal server error"))

	// Act
	err := commentService.Create(comment)

	// Assert
	assert.Error(t, err)
}

// TestCommentService_Delete tests the Delete method of the CommentService
func TestCommentService_Delete(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	commentRepository := mock_interfaces.NewMockCommentRepository(ctrl)
	userRepository := mock_interfaces.NewMockUserRepository(ctrl)
	commentService := NewCommentService(commentRepository, userRepository)

	// Mock repository response
	commentRepository.EXPECT().Delete("some_id").Return(nil)

	// Act
	err := commentService.Delete("some_id")

	// Assert
	assert.Nil(t, err)
}

// TestCommentService_Delete_Fail tests the Delete method of the CommentService when repository fail
func TestCommentService_Delete_Fail(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	commentRepository := mock_interfaces.NewMockCommentRepository(ctrl)
	userRepository := mock_interfaces.NewMockUserRepository(ctrl)
	commentService := NewCommentService(commentRepository, userRepository)

	// Mock repository response
	commentRepository.EXPECT().Delete("some_id").Return(errors.NewAppError(500, "Internal server error"))

	// Act
	err := commentService.Delete("some_id")

	// Assert
	assert.Error(t, err)
}

// TestCommentService_GetNoteComments tests the GetNoteComments method of the CommentService
func TestCommentService_GetNoteComments(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	commentRepository := mock_interfaces.NewMockCommentRepository(ctrl)
	userRepository := mock_interfaces.NewMockUserRepository(ctrl)
	commentService := NewCommentService(commentRepository, userRepository)

	comment1 := domain.Comment{
		Id:       "some_id1",
		UserId:   "user_id1",
		NoteId:    "note_id1",
		Content:   "content1",
		CreatedAt: "created_at1",
	}

	comment2 := domain.Comment{
		Id:       "some_id2",
		UserId:   "user_id2",
		NoteId:    "note_id2",
		Content:   "content2",
		CreatedAt: "created_at2",
	}

	// Mock repository response
	commentRepository.EXPECT().GetNoteComments("some_id").Return([]domain.Comment{comment1, comment2}, nil)

	// Act
	notes, err := commentService.GetNoteComments("some_id")

	noteComment1 := domain.NoteComment{
		Id:       "some_id1",
		Author:   "user_id1",
		Content:   "content1",
		CreatedAt: "created_at1",
	}

	noteComment2 := domain.NoteComment{
		Id:       "some_id2",
		Author:   "user_id2",
		Content:   "content2",
		CreatedAt: "created_at2",
	}

	// Assert
	assert.Equal(t, notes, []domain.NoteComment{noteComment1, noteComment2})
	assert.Nil(t, err)
}

// TestCommentService_GetNoteComments tests the GetNoteComments method of the CommentService when a note doesn't exist
func TestCommentService_GetNoteComments_FailWhenNoteNotExists(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	commentRepository := mock_interfaces.NewMockCommentRepository(ctrl)
	userRepository := mock_interfaces.NewMockUserRepository(ctrl)
	commentService := NewCommentService(commentRepository, userRepository)

	// Mock repository response
	commentRepository.EXPECT().GetNoteComments("some_id").Return(nil, errors.NewAppError(500, "Internal server error"))

	// Act
	notes, err := commentService.GetNoteComments("some_id")

	// Assert
	assert.Nil(t, notes)
	assert.Error(t, err)
}

// TestCommentService_CountCommentByIdNoteController tests the CountCommentByIdNoteController method of the CommentService
func TestCommentService_CountCommentByIdNoteController(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	commentRepository := mock_interfaces.NewMockCommentRepository(ctrl)
	userRepository := mock_interfaces.NewMockUserRepository(ctrl)
	commentService := NewCommentService(commentRepository, userRepository)

	comment1 := domain.Comment{
		Id:       "some_id1",
		UserId:   "user_id1",
		NoteId:    "note_id1",
		Content:   "content1",
		CreatedAt: "created_at1",
	}

	comment2 := domain.Comment{
		Id:       "some_id2",
		UserId:   "user_id2",
		NoteId:    "note_id2",
		Content:   "content2",
		CreatedAt: "created_at2",
	}

	// Mock repository response
	commentRepository.EXPECT().GetNoteComments("some_id").Return([]domain.Comment{comment1, comment2}, nil)

	// Act
	numberComments, err := commentService.CountCommentByIdNoteController("some_id")

	// Assert
	assert.Equal(t, numberComments, 2)
	assert.Nil(t, err)
}

// TestCommentService_CountCommentByIdNoteController_FailWhenNoteNotExists tests the CountCommentByIdNoteController method of the CommentService when a note doesn't exist
func TestCommentService_CountCommentByIdNoteController_FailWhenNoteNotExists(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	commentRepository := mock_interfaces.NewMockCommentRepository(ctrl)
	userRepository := mock_interfaces.NewMockUserRepository(ctrl)
	commentService := NewCommentService(commentRepository, userRepository)

	// Mock repository response
	commentRepository.EXPECT().GetNoteComments("some_id").Return(nil, errors.NewAppError(500, "Internal server error"))

	// Act
	numberComments, err := commentService.CountCommentByIdNoteController("some_id")

	// Assert
	assert.Equal(t, numberComments, 0)
	assert.Error(t, err)
}