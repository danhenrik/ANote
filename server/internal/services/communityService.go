package services

import (
	"anote/internal/domain"
	"anote/internal/errors"
	"anote/internal/helpers"
	IRepo "anote/internal/interfaces/repositories"
)

type CommunityService struct {
	communityRepo IRepo.CommunityRepository
}

func NewCommunityService(repo IRepo.CommunityRepository) CommunityService {
	return CommunityService{communityRepo: repo}
}

func (this CommunityService) Create(community *domain.Community) *errors.AppError {
	community.Id = helpers.NewUUID()
	err := this.communityRepo.Create(community)
	if err != nil {
		return err
	}
	return nil
}

func (this CommunityService) GetAll() ([]domain.Community, *errors.AppError) {
	communities, err := this.communityRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return communities, nil
}

func (this CommunityService) Join(id string, userID string) *errors.AppError {
	err := this.communityRepo.AddMember(id, userID)
	if err != nil {
		return err
	}
	return nil
}

func (this CommunityService) Leave(id string, userID string) *errors.AppError {
	err := this.communityRepo.RemoveMember(id, userID)
	if err != nil {
		return err
	}
	return nil
}

func (this CommunityService) Delete(id string) *errors.AppError {
	err := this.communityRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
