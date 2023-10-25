package services

import (
	"anote/internal/domain"
	"anote/internal/errors"
	"anote/internal/helpers"
	IRepo "anote/internal/interfaces/repositories"
	"anote/internal/storage/es"
	"anote/internal/viewmodels"
	"fmt"
	"log"
	"slices"
)

type NoteService struct {
	userRepository      IRepo.UserRepository
	communityRepository IRepo.CommunityRepository
	noteRepository      IRepo.NoteRepository
	noteTagRepository   IRepo.NoteTagRepository
}

func NewNoteService(
	userRepo IRepo.UserRepository,
	communityRepo IRepo.CommunityRepository,
	noteRepo IRepo.NoteRepository,
	tagRepo IRepo.NoteTagRepository,
) NoteService {
	return NoteService{
		userRepository:      userRepo,
		noteRepository:      noteRepo,
		noteTagRepository:   tagRepo,
		communityRepository: communityRepo,
	}
}

func (this NoteService) Create(
	note *domain.Note,
	tagIDs []string,
	communityIDs []string,
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

	for _, communityID := range communityIDs {
		community, err := this.communityRepository.GetById(communityID)
		if community == nil {
			return "", errors.NewAppError(400, "Invalid community provided")
		}
		if err != nil {
			log.Println("[NoteService] Error on get community by id:", err)
			return "", err
		}
		isMember, err := this.communityRepository.CheckMember(communityID, user.Id)
		if err != nil {
			log.Println("[NoteService] Error on check member:", err)
			return "", errors.NewAppError(400, "Error while checking if member is part of community")
		}
		if !isMember {
			return "", errors.NewAppError(400, "User not member of community")
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
		return "", errors.NewAppError(200, "Note created but couldn't save tags")
	}
	err = this.noteRepository.AddCommunities(note.Id, communityIDs)
	if err != nil {
		log.Println("[NoteService] Error on add communities:", err)
		return "", errors.NewAppError(200, "Note created but couldn't save communities")
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

	communities, err := this.communityRepository.GetByNoteId(note.Id)
	if err != nil {
		log.Println("[NoteService] Error on get note communities:", err)
		return nil, err
	}

	fnote := domain.FullNote{
		Id:          note.Id,
		Title:       note.Title,
		Content:     note.Content,
		AuthorID:    note.AuthorID,
		CreatedAt:   note.CreatedAt,
		UpdatedAt:   note.UpdatedAt,
		Tags:        tags,
		Communities: communities,
	}
	return &fnote, nil
}

func (this NoteService) GetByCommunityId(id string) ([]domain.FullNote, *errors.AppError) {
	notes, err := this.noteRepository.GetByCommunityID(id)
	if notes == nil {
		return nil, errors.NewAppError(404, "Note not found")
	}
	if err != nil {
		log.Println("[NoteService] Error on get note:", err)
		return nil, err
	}

	fnotes := []domain.FullNote{}
	for _, note := range notes {
		tags, err := this.noteTagRepository.GetByNoteId(note.Id)
		if err != nil {
			log.Println("[NoteService] Error on get note tags:", err)
			return nil, err
		}

		communities, err := this.communityRepository.GetByNoteId(note.Id)
		if err != nil {
			log.Println("[NoteService] Error on get note communities:", err)
			return nil, err
		}

		fnote := domain.FullNote{
			Id:          note.Id,
			Title:       note.Title,
			Content:     note.Content,
			AuthorID:    note.AuthorID,
			CreatedAt:   note.CreatedAt,
			UpdatedAt:   note.UpdatedAt,
			Tags:        tags,
			Communities: communities,
		}
		fnotes = append(fnotes, fnote)
	}
	return fnotes, nil
}

func (this NoteService) GetByAuthorId(id string) ([]domain.FullNote, *errors.AppError) {
	notes, err := this.noteRepository.GetByAuthorID(id)
	if notes == nil {
		return nil, errors.NewAppError(404, "Note not found")
	}
	if err != nil {
		log.Println("[NoteService] Error on get note:", err)
		return nil, err
	}

	fnotes := []domain.FullNote{}
	for _, note := range notes {
		tags, err := this.noteTagRepository.GetByNoteId(note.Id)
		if err != nil {
			log.Println("[NoteService] Error on get note tags:", err)
			return nil, err
		}

		communities, err := this.communityRepository.GetByNoteId(note.Id)
		if err != nil {
			log.Println("[NoteService] Error on get note communities:", err)
			return nil, err
		}

		fnote := domain.FullNote{
			Id:          note.Id,
			Title:       note.Title,
			Content:     note.Content,
			AuthorID:    note.AuthorID,
			CreatedAt:   note.CreatedAt,
			UpdatedAt:   note.UpdatedAt,
			Tags:        tags,
			Communities: communities,
		}
		fnotes = append(fnotes, fnote)
	}
	return fnotes, nil
}

var convertMap = map[string]string{
	"id":             "id.keyword",
	"title":          "title.keyword",
	"content":        "content.keyword",
	"published_date": "published_date",
	"updated_date":   "updated_date",
	"author":         "author.keyword",
	"communities":    "communities.name.keyword",
	"tags":           "tags.name.keyword",
	"likes":          "likes_count",
	"comment":        "comments_count",
}

func (this NoteService) Search(
	page int,
	size int,
	title string,
	content string,
	authorID string,
	tags []string,
	communities []string,
	createdAt [2]string,
	sortOrder string,
	sortField string,
) ([]domain.FilteredNote, *errors.AppError) {
	qb := es.NewNoteQueryBuilder()
	if page <= 0 || size <= 0 {
		return nil, errors.NewAppError(400, "Invalid pagination parameters")
	}
	qb.AddPagination(page, size)

	if sortField != "" && sortOrder != "" {
		qb.AddSort(convertMap[sortField], sortOrder)
	}

	if len(tags) > 0 {
		qb.AddTagsQuery(tags)
	}

	if authorID != "" {
		qb.AddAuthorQuery(authorID)
	}

	if title != "" {
		qb.AddTitleQuery(title)
	}

	if content != "" {
		qb.AddContentQuery(content)
	}

	if createdAt[0] != "" && createdAt[1] == "" {
		ok := helpers.ValidateDate(createdAt[0])
		if !ok {
			return nil, errors.NewAppError(400, "Invalid time format")
		}
		qb.AddCreatedAtMatchQuery(createdAt[0])
	} else if createdAt[0] != "" && createdAt[1] != "" {
		ok := helpers.ValidateDate(createdAt[0])
		if !ok {
			return nil, errors.NewAppError(400, "Invalid upper limit time format")
		}
		ok = helpers.ValidateDate(createdAt[1])
		if !ok {
			return nil, errors.NewAppError(400, "Invalid upper limit time format")
		}
		qb.AddCreatedAtRangeQuery(createdAt[0], createdAt[1])
	}

	if len(communities) > 0 {
		qb.AddCommunitiesQuery(communities)
	}

	notes, err := qb.Query()
	if err != nil {
		log.Println("[NoteService] Error on query notes:", err)
		return nil, err
	}
	return notes, nil
}

func (this NoteService) GetFeed(
	page int,
	size int,
	authorID string,
	communityIds []string,
	sortOrder string,
	sortField string,
) ([]domain.FilteredNote, *errors.AppError) {
	qb := es.NewNoteQueryBuilder()
	if page <= 0 || size <= 0 {
		return nil, errors.NewAppError(400, "Invalid pagination parameters")
	}
	qb.AddPagination(page, size)

	if sortField != "" && sortOrder != "" {
		qb.AddSort(convertMap[sortField], sortOrder)
	}

	// notes in communities
	if len(communityIds) > 0 {
		qb.AddCommunityIdsQuery(communityIds)
	}

	// or private (created by author)
	qb.FinishShould().
		AddAuthorQuery(authorID)

	notes, err := qb.Query()
	if err != nil {
		log.Println("[NoteService] Error on query notes:", err)
		return nil, err
	}
	return notes, nil
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
	communities, err := this.communityRepository.GetByNoteId(noteVM.Id)
	if err != nil {
		log.Println("[NoteService] Error on get note communities:", err)
		return err
	}
	// validate communities
	for _, communityID := range noteVM.RemoveCommunities {
		community, err := this.communityRepository.GetById(communityID)
		if community == nil {
			return errors.NewAppError(400, fmt.Sprintf("Invalid remove community provided w/ id %s not found", communityID))
		}
		if err != nil {
			log.Println("[NoteService] Error on get remove community by id:", err)
			return err
		}
		if !slices.Contains(communities, *community) {
			return errors.NewAppError(400, fmt.Sprintf("Remove community w/id %s not present in note communities", communityID))
		}
	}
	for _, communityID := range noteVM.AddCommunities {
		community, err := this.communityRepository.GetById(communityID)
		if community == nil {
			return errors.NewAppError(400, fmt.Sprintf("Invalid add community provided w/ id %s", communityID))
		}
		if err != nil {
			log.Println("[NoteService] Error on get add community by id:", err)
			return err
		}
		if slices.Contains(communities, *community) {
			return errors.NewAppError(400, fmt.Sprintf("Add community w/id %s already present in note communities", communityID))
		}
		userIsMember, err := this.communityRepository.CheckMember(communityID, requestUserId)
		if err != nil {
			log.Println("[NoteService] Error on check member:", err)
			return errors.NewAppError(400, "Error while checking if member is part of community:"+communityID)
		}
		if !userIsMember {
			return errors.NewAppError(400, "User not member of community: "+communityID)
		}
	}

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
	if len(noteVM.AddCommunities) > 0 {
		err = this.noteRepository.AddCommunities(note.Id, noteVM.AddCommunities)
		if err != nil {
			log.Println("[NoteService] Error on add communities:", err)
			return err
		}
	}
	if len(noteVM.RemoveCommunities) > 0 {
		err = this.noteRepository.RemoveCommunities(note.Id, noteVM.RemoveCommunities)
		if err != nil {
			log.Println("[NoteService] Error on remove communities:", err)
			return err
		}
	}
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

func (this NoteService) GetAll() ([]domain.FullNote, *errors.AppError) {
	notes, err := this.noteRepository.GetAll()
	if err != nil {
		log.Println("[NoteService] Error on get note:", err)
		return nil, err
	}

	var fnote []domain.FullNote
	for _, note := range notes {
		tags, errTag := this.noteTagRepository.GetByNoteId(note.Id)
		if errTag != nil {
			log.Println("[NoteService] Error on get note tags:", errTag)
			return nil, errTag
		}

		communities, errComm := this.communityRepository.GetByNoteId(note.Id)
		if errComm != nil {
			log.Println("[NoteService] Error on get note communities:", errComm)
			return nil, errComm
		}

		fnote = append(fnote, domain.FullNote{
			Id:          note.Id,
			Title:       note.Title,
			Content:     note.Content,
			AuthorID:    note.AuthorID,
			CreatedAt:   note.CreatedAt,
			UpdatedAt:   note.UpdatedAt,
			Tags:        tags,
			Communities: communities,
		})
	}

	return fnote, nil
}
