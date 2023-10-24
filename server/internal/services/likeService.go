package services

import (
	"anote/internal/domain"
	"anote/internal/errors"
	"anote/internal/helpers"
	IRepo "anote/internal/interfaces/repositories"
)

type LikeService struct {
	likeRepository IRepo.LikeRepository
}

func NewLikeService(likeRepository IRepo.LikeRepository) LikeService {
	return LikeService{ likeRepository: likeRepository }
}

func (this LikeService) Create(like *domain.Like) *errors.AppError {
	like.Id = helpers.NewUUID()
	err := this.likeRepository.Create(like)
	if err != nil {
		return err
	}
	return nil
}

func (this LikeService) Delete(idUser string, idNote string) *errors.AppError {
	err := this.likeRepository.Delete(idUser, idNote)
	if err != nil {
		return err
	}
	return nil
}