package services

import (
	"anote/internal/domain"
	"anote/internal/errors"
	"anote/internal/helpers"
	IRepo "anote/internal/interfaces/repositories"
	"log"
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
		log.Println("[LikeService] Error on create like:", err)
		return err
	}

	return nil
}

func (this LikeService) Delete(idUser string, idNote string) *errors.AppError {
	err := this.likeRepository.Delete(idUser, idNote)

	if err != nil {
		log.Println("[LikeService] Error on delete like:", err)
		return err
	}

	return nil
}

func (this LikeService) GetByIdUserAndIdNote(idUser string, idNote string) (*domain.Like, *errors.AppError) {
	like, err := this.likeRepository.GetByIdUserAndIdNote(idUser, idNote)

	if err != nil {
		log.Println("[LikeService] Error on get like:", err)
		return nil, err
	}

	return like, nil
}

func (this LikeService) CountLikeByIdNoteController(idNote string) (int, *errors.AppError) {
	likes, err := this.likeRepository.GetByIdNote(idNote)

	if err != nil {
		log.Println("[LikeService] Error on count like:", err)
		return 0, err
	}

	var numberLikes int = 0

	for range likes {
		numberLikes++
	}

	return numberLikes, nil
}