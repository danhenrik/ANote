package services

import (
	"anote/internal/domain"
	"anote/internal/errors"
	"anote/internal/helpers"
	IRepo "anote/internal/interfaces/repositories"
	"log"
)

type NoteService struct {
	userRepository    IRepo.UserRepository
	noteRepository    IRepo.NoteRepository
	noteTagRepository IRepo.NoteTagRepository
}

func NewNoteService(
	userRepo IRepo.UserRepository,
	noteRepo IRepo.NoteRepository,
	tagRepo IRepo.NoteTagRepository,
) NoteService {
	return NoteService{
		userRepository:    userRepo,
		noteRepository:    noteRepo,
		noteTagRepository: tagRepo,
	}
}

func (this NoteService) Create(
	note *domain.Note,
	tagIDs []string,
) (string, *errors.AppError) {
	user, _ := this.userRepository.GetByUsername(note.AuthorID)
	if user == nil {
		return "", errors.NewAppError(400, "User logged in not found")
	}

	for _, tagID := range tagIDs {
		tag, _ := this.noteTagRepository.GetById(tagID)
		if tag == nil {
			return "", errors.NewAppError(400, "Invalid tag provided")
		}
	}

	note.Id = helpers.NewUUID()
	err := this.noteRepository.Create(note)
	if err != nil {
		log.Println("[NoteService] Error on create note:", err)
		return "", err
	}
	return note.Id, nil
}
