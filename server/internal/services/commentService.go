package services

import (
	"anote/internal/domain"
	"anote/internal/errors"
	"anote/internal/helpers"
	IRepo "anote/internal/interfaces/repositories"
	"log"
)

type CommentService struct {
	commentRepository IRepo.CommentRepository
	userRepository    IRepo.UserRepository
}

func NewCommentService(
	commentRepository IRepo.CommentRepository,
	userRepo 		  IRepo.UserRepository,
) CommentService {
	return CommentService{
		commentRepository: commentRepository,
		userRepository:    userRepo,
	}
}

func (this CommentService) Create(comment *domain.Comment) *errors.AppError {
	comment.Id = helpers.NewUUID()
	err := this.commentRepository.Create(comment)

	if err != nil {
		log.Println("[CommentService] Error on create comment:", err)
		return err
	}

	return nil
}

func (this CommentService) Delete(idComment string) *errors.AppError {
	err := this.commentRepository.Delete(idComment)

	if err != nil {
		log.Println("[CommentService] Error on delete comment:", err)
		return err
	}

	return nil
}

func (this CommentService) GetNoteComments(idNote string) ([]domain.NoteComment, *errors.AppError) {
	comments, err := this.commentRepository.GetNoteComments(idNote)

	if err != nil {
		log.Println("[CommentService] Error on get note comments:", err)
		return nil, err
	}

	var noteComments []domain.NoteComment

	for _, comment := range comments {
		noteComments = append(noteComments, domain.NoteComment{
			Id:        comment.Id,
			Author:    comment.UserId,
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt,
		})
	}

	return noteComments, nil
}

func (this CommentService) CountCommentByIdNoteController(idNote string) (int, *errors.AppError) {
	comments, err := this.commentRepository.GetNoteComments(idNote)

	if err != nil {
		log.Println("[CommentService] Error on count comment:", err)
		return 0, err
	}

	var numberComments int = 0

	for range comments {
		numberComments++
	}

	return numberComments, nil
}