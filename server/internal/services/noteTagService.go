package services

import (
	"anote/internal/domain"
	"anote/internal/errors"
	"anote/internal/helpers"
	IRepo "anote/internal/interfaces/repositories"
)

type NoteTagService struct {
	noteTagRepo IRepo.NoteTagRepository
}

func NewNoteTagService(repo IRepo.NoteTagRepository) NoteTagService {
	return NoteTagService{noteTagRepo: repo}
}

func (this NoteTagService) Create(tag *domain.NoteTag) *errors.AppError {
	tag.Id = helpers.NewUUID()
	err := this.noteTagRepo.Create(tag)
	if err != nil {
		return err
	}
	return nil
}
