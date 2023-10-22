package services

import (
	"anote/internal/domain"
	"anote/internal/errors"
	"anote/internal/helpers"
	IRepo "anote/internal/interfaces/repositories"
	"anote/internal/viewmodels"
	"fmt"
	"log"
	"slices"
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
		tag, err := this.noteTagRepository.GetById(tagID)
		if tag == nil {
			return "", errors.NewAppError(400, "Invalid tag provided")
		}
		if err != nil {
			log.Println("[NoteService] Error on get tag by id:", err)
			return "", err
		}
	}

	note.Id = helpers.NewUUID()
	err := this.noteRepository.Create(note)
	if err != nil {
		log.Println("[NoteService] Error on create note:", err)
		return "", err
	}

	err = this.noteRepository.AddTags(note.Id, tagIDs)
	if err != nil {
		log.Println("[NoteService] Error on add tags:", err)
		return "", errors.NewAppError(500, "Error on add tags, note created")
	}
	return note.Id, nil
}

func (this NoteService) GetById(id string) (*domain.FullNote, *errors.AppError) {
	note, err := this.noteRepository.GetByID(id)
	if note == nil {
		return nil, errors.NewAppError(404, "Note not found")
	}
	if err != nil {
		log.Println("[NoteService] Error on get note:", err)
		return nil, err
	}

	tags, err := this.noteTagRepository.GetByNoteId(note.Id)
	if err != nil {
		log.Println("[NoteService] Error on get note tags:", err)
		return nil, err
	}

	fnote := domain.FullNote{
		Id:        		note.Id,
		Title:     		note.Title,
		Content:   		note.Content,
		AuthorID:  		note.AuthorID,
		CreatedAt: 		note.CreatedAt,
		UpdatedAt:	 	note.UpdatedAt,
		LikeCount: 		note.LikeCount,
		CommentCount: note.CommentCount,
		Tags:      tags,
		// TODO: Insert communities
	}
	return &fnote, nil
}

func (this NoteService) Update(requestUserId string, noteVM viewmodels.UpdateNoteVM) *errors.AppError {
	// get note
	note, err := this.noteRepository.GetByID(noteVM.Id)
	if note == nil {
		return errors.NewAppError(404, "Note not found")
	}
	if err != nil {
		log.Println("[NoteService] Error on get note:", err)
		return err
	}

	// check if user is author
	if note.AuthorID != requestUserId {
		return errors.NewAppError(400, "User not author of note")
	}

	// get tags
	tags, err := this.noteTagRepository.GetByNoteId(noteVM.Id)
	if err != nil {
		log.Println("[NoteService] Error on get note tags:", err)
		return err
	}
	// validate tags
	for _, tagID := range noteVM.RemoveTags {
		tag, err := this.noteTagRepository.GetById(tagID)
		if tag == nil {
			return errors.NewAppError(400, fmt.Sprintf("Invalid remove tag provided w/ id %s not found", tagID))
		}
		if err != nil {
			log.Println("[NoteService] Error on get remove tag by id:", err)
			return err
		}
		if !slices.Contains(tags, *tag) {
			return errors.NewAppError(400, fmt.Sprintf("Remove tag w/id %s not present in note tags", tagID))
		}
	}
	for _, tagID := range noteVM.AddTags {
		tag, err := this.noteTagRepository.GetById(tagID)
		if tag == nil {
			return errors.NewAppError(400, fmt.Sprintf("Invalid add tag provided w/ id %s", tagID))
		}
		if err != nil {
			log.Println("[NoteService] Error on get add tag by id:", err)
			return err
		}
		if slices.Contains(tags, *tag) {
			return errors.NewAppError(400, fmt.Sprintf("Add tag w/id %s already present in note tags", tagID))
		}
	}

	// get communities

	// check if user is member of communities

	// update note
	if noteVM.Title != "" {
		note.Title = noteVM.Title
	}
	if noteVM.Content != "" {
		note.Content = noteVM.Content
	}
	err = this.noteRepository.Update(note)
	if err != nil {
		log.Println("[NoteService] Error on update note:", err)
		return err
	}

	// update tags
	if len(noteVM.AddTags) > 0 {
		err = this.noteRepository.AddTags(note.Id, noteVM.AddTags)
		if err != nil {
			log.Println("[NoteService] Error on add tags:", err)
			return err
		}
	}
	if len(noteVM.RemoveTags) > 0 {
		err = this.noteRepository.RemoveTags(note.Id, noteVM.RemoveTags)
		if err != nil {
			log.Println("[NoteService] Error on remove tags:", err)
			return err
		}
	}

	// update communities
	return nil
}

func (this NoteService) Delete(id string) *errors.AppError {
	err := this.noteRepository.Delete(id)
	if err != nil {
		log.Println("[NoteService] Error on delete note:", err)
		return err
	}
	return nil
}

func (this NoteService) GetAll() ([]domain.FullNoteList, *errors.AppError) {
	notes, err := this.noteRepository.GetAll()
	if err != nil {
		log.Println("[NoteService] Error on get note:", err)
		return nil, err
	}

	var fnote []domain.FullNoteList
	for _, note := range notes {
		tags, errTag := this.noteTagRepository.GetByNoteId(note.Id)
		user, errUser := this.userRepository.GetByUsername(note.AuthorID)
		if errTag != nil {
			log.Println("[NoteService] Error on get note tags:", errTag)
			return nil, errTag
		} else if errUser != nil {
			log.Println("[NoteService] Error on get user note:", errUser)
			return nil, errUser
		}
	
		var tagsDescription []string
		for _, tag := range tags {
			tagsDescription = append(tagsDescription, tag.Name)
		}

		fnote = append(fnote, domain.FullNoteList{
			Id:        		note.Id,
			Title:     		note.Title,
			Content:   		note.Content,
			AuthorID:  		note.AuthorID,
			Author:				user.Email,
			PublishedDate: 		note.CreatedAt,
			UpdatedDate: 		note.UpdatedAt,
			LikesCount: 		note.LikeCount,
			CommentCount: note.CommentCount,
			Tags:      		tagsDescription,
			// TODO: Insert communities
		})
	}

	return fnote, nil
}