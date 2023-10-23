package interfaces

import (
	"anote/internal/domain"
	"anote/internal/errors"
)

type CommunityRepository interface {
	Create(community *domain.Community) *errors.AppError
	GetAll() ([]domain.Community, *errors.AppError)
	GetById(id string) (*domain.Community, *errors.AppError)
	GetByNoteId(noteId string) ([]domain.Community, *errors.AppError)
	GetByUserId(userId string) ([]domain.Community, *errors.AppError)
	AddMember(communityId string, userId string) *errors.AppError
	RemoveMember(communityId string, userId string) *errors.AppError
	GetMembers(communityId string) ([]domain.User, *errors.AppError)
	CheckMember(communityId string, userId string) (bool, *errors.AppError)
	Delete(id string) *errors.AppError
}
