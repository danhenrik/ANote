// services/user_service_test.go
package services

import (
	"anote/internal/domain"
	"anote/internal/errors"
	"anote/internal/helpers"
	mock_interfaces "anote/internal/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func StringPointer(s string) *string {
    return &s
}

// TestUserService_Create tests the Create method of the UserService
func TestUserService_Create(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepository := mock_interfaces.NewMockUserRepository(ctrl)
	userService := NewUserService(userRepository)

	user := &domain.User{
		Id:        "some_id",
		Email:     "test@example.com",
		Password:  StringPointer("password123"),
		Google_id: StringPointer(""),
		CreatedAt: "some_timestamp",
		Avatar:    StringPointer("test.png"),
	}
	// Mock repository response
	userRepository.EXPECT().Create(user).Return(errors.NewAppError(500, "Internal server error"))

	// Act
	err := userService.Create(user)

	// Assert
	assert.Error(t, err)
}

func TestUserService_UpdatePassword(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    userRepository := mock_interfaces.NewMockUserRepository(ctrl)
    userService := NewUserService(userRepository)

    username := "user123"
    oldPassword := "oldPassword"
    newPassword := "newPassword"

    hashedOldPassword, _ := helpers.Hash(oldPassword)
	userRepository.EXPECT().GetUserWithPassword(username).Return(&domain.User{
		Id:        "some_id",
		Email:     "test@example.com",
		Password:  StringPointer(hashedOldPassword),
		Google_id: StringPointer(""),
		CreatedAt: "some_timestamp",
		Avatar:    StringPointer("test.png"),
	}, nil)

	userRepository.EXPECT().UpdatePassword(username, gomock.Any()).Return(nil)

    // Act
    err := userService.UpdatePassword(username, oldPassword, newPassword)
	
    // Assert
    assert.Nil(t, err)
}

/*
func TestUserService_SaveAvatar(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepository := mock_interfaces.NewMockUserRepository(ctrl)
	userService := NewUserService(userRepository)

	userID := "user123"
	filename := "avatar.png"

	t.Log("UserID:", userID)
	t.Log("Filename:", filename)

	userRepository.EXPECT().GetByUsername(userID).Return(&domain.User{
		Id:        "some_id",
		Email:     "test@example.com",
		Password:  StringPointer("password123"),
		Google_id: StringPointer(""),
		CreatedAt: "some_timestamp",
		Avatar:    StringPointer("test.png"),
	}, nil).AnyTimes()
	userRepository.EXPECT().SetAvatar(userID, filename).Return(nil).AnyTimes()

	err := userService.SaveAvatar(userID, filename)

	assert.NoError(t, err)
}
*/
