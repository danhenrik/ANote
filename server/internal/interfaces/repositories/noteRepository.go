package interfaces

import (
	"anote/internal/domain"
	"anote/internal/errors"
)

type NoteRepository interface {
	Create(note *domain.Note) *errors.AppError
	AddTags(noteId string, tagIds []string) *errors.AppError
	RemoveTags(noteId string, tagIds []string) *errors.AppError
	AddCommunities(noteId string, communityIds []string) *errors.AppError
	RemoveCommunities(noteId string, communityIds []string) *errors.AppError
	GetAll() ([]domain.Note, *errors.AppError)
	GetByID(id string) (*domain.Note, *errors.AppError)
	GetByAuthorID(userID string) ([]domain.Note, *errors.AppError)
	GetByCommunityID(communityID string) ([]domain.Note, *errors.AppError)
	Update(note *domain.Note) *errors.AppError
	Delete(id string) *errors.AppError
}
