package services

import (
	"anote/internal/domain"
	"anote/internal/errors"
	"anote/internal/helpers"
	IRepo "anote/internal/interfaces/repositories"
	"io"
	"os"
	"strings"
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

func (this CommunityService) GetByUserId(id string) ([]domain.Community, *errors.AppError) {
	communities, err := this.communityRepo.GetByUserId(id)
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

func (this CommunityService) DeleteBackground(communityId string) *errors.AppError {
	community, err := this.communityRepo.GetById(communityId)
	if err != nil {
		return err
	}

	delErr := os.Remove("internal/assets/" + *community.Background)
	if delErr != nil {
		return errors.NewAppError(500, "Internal server error")
	}

	err = this.communityRepo.SetBackground(communityId, "")

	return nil
}

func (this CommunityService) SaveBackground(communityId string, filename string) *errors.AppError {
	splitted := strings.Split(filename, ".")
	ext := splitted[len(splitted)-1]
	backgroundFilename := communityId + "." + ext

	// create permanent file
	file, err := os.OpenFile("internal/assets/"+backgroundFilename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return errors.NewAppError(500, "Internal server error")
	}
	defer file.Close()

	// open temporary file
	tempFile, err := os.OpenFile("internal/tmp/"+filename, os.O_RDONLY, 0777)
	if err != nil {
		return errors.NewAppError(500, "Internal server error")
	}
	defer tempFile.Close()

	// move (copy + delete)
	_, err = io.Copy(file, tempFile)
	if err != nil {
		return errors.NewAppError(500, "Internal server error")
	}
	os.Remove("internal/tmp/" + filename)

	this.communityRepo.SetBackground(communityId, backgroundFilename)
	return nil
}
