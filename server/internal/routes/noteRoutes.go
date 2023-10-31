package routes

import (
	httpAdapter "anote/cmd/interfaces"
	"anote/internal/container"
	"anote/internal/viewmodels"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func CreateNoteController(request httpAdapter.Request) httpAdapter.Response {
	var noteVM viewmodels.CreateNoteVM
	if err := json.Unmarshal([]byte(request.Body), &noteVM); err != nil {
		log.Println("[NoteController] Error on create note unmarshal:", err)
		return httpAdapter.NewErrorResponse(400, "Invalid content-type")
	}

	note := noteVM.ToDomainNote()
	note.AuthorID = request.User.ID

	noteId, err := container.NoteService.Create(&note, noteVM.Tags, noteVM.Communities)
	if err != nil {
		log.Println("[NoteController] Error on create note:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}
	return httpAdapter.NewSuccessResponse(201, map[string]string{"id": noteId})
}

func UpdateNoteController(request httpAdapter.Request) httpAdapter.Response {
	var noteVM viewmodels.UpdateNoteVM
	if err := json.Unmarshal([]byte(request.Body), &noteVM); err != nil {
		log.Println("[NoteController] Error on update note unmarshal:", err)
		return httpAdapter.NewErrorResponse(400, "Invalid content-type")
	}

	err := container.NoteService.Update(request.User.ID, noteVM)
	if err != nil {
		log.Println("[NoteController] Error on update note:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}
	return httpAdapter.NewNoContentResponse()
}

func SearchNoteController(request httpAdapter.Request) httpAdapter.Response {
	page, _ := request.GetSingleQuery("page")
	pageInt, convErr := strconv.Atoi(page)
	if convErr != nil {
		pageInt = 0
	}
	size, _ := request.GetSingleQuery("size")
	sizeInt, convErr := strconv.Atoi(size)
	if convErr != nil {
		sizeInt = 0
	}
	title, _ := request.GetSingleQuery("title")
	content, _ := request.GetSingleQuery("content")
	tags, _ := request.GetQuerySlice("tags")
	author, _ := request.GetSingleQuery("author")
	communities, _ := request.GetQuerySlice("communities")
	published_date, _ := request.GetSingleQuery("date")
	from_published_date, _ := request.GetSingleQuery("from_date")
	to_published_date, _ := request.GetSingleQuery("to_date")
	sortOrder, _ := request.GetSingleQuery("sort_order")
	if sortOrder == "" {
		sortOrder = "desc"
	}
	sortField, _ := request.GetSingleQuery("sort_by")

	createdAt := [2]string{"", ""}
	if published_date != "" {
		createdAt[0] = published_date
	} else if from_published_date != "" && to_published_date != "" {
		createdAt[0] = from_published_date
		createdAt[1] = to_published_date
	}

	notes, err := container.NoteService.Search(
		pageInt,
		sizeInt,
		title,
		content,
		author,
		tags,
		communities,
		createdAt,
		sortOrder,
		sortField,
	)
	if err != nil {
		log.Println("[NoteController] Error on search note:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}
	return httpAdapter.NewSuccessResponse(200, notes)
}

func GetNoteFeedController(request httpAdapter.Request) httpAdapter.Response {
	page, _ := request.GetSingleQuery("page")
	pageInt, convErr := strconv.Atoi(page)
	if convErr != nil {
		pageInt = 0
	}
	size, _ := request.GetSingleQuery("size")
	sizeInt, convErr := strconv.Atoi(size)
	if convErr != nil {
		sizeInt = 0
	}
	sortOrder, _ := request.GetSingleQuery("sort_order")
	if sortOrder == "" {
		sortOrder = "desc"
	}
	sortField, _ := request.GetSingleQuery("sort_by")

	communities, err := container.CommunityService.GetByUserId(request.User.ID)
	if err != nil {
		log.Println("[NoteController] Error on get note feed:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}
	communityIds := make([]string, len(communities))
	for i, id := range communities {
		communityIds[i] = id.Id
	}
	notes, err := container.NoteService.GetFeed(
		pageInt,
		sizeInt,
		request.User.ID,
		communityIds,
		sortOrder,
		sortField,
	)
	if err != nil {
		log.Println("[NoteController] Error on get note feed:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}
	return httpAdapter.NewSuccessResponse(200, notes)
}

func GetNoteByIDController(request httpAdapter.Request) httpAdapter.Response {
	id, ok := request.GetSingleParam("id")
	if !ok {
		log.Println("[NoteController] No id provided on get:")
		return httpAdapter.NewErrorResponse(400, "Invalid note id")
	}

	note, err := container.NoteService.GetById(id)
	if note == nil {
		return httpAdapter.NewErrorResponse(404, "Note not found")
	}
	if err != nil {
		log.Println("[NoteController] Error on get note:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}

	noteVM := viewmodels.NoteVM{}.FromDomain(*note)
	return httpAdapter.NewSuccessResponse(200, noteVM)
}

func GetNoteByAuthorIDController(request httpAdapter.Request) httpAdapter.Response {
	id, ok := request.GetSingleParam("id")
	if !ok {
		log.Println("[NoteController] No id provided on get:")
		return httpAdapter.NewErrorResponse(400, "Invalid author id")
	}

	notes, err := container.NoteService.GetByAuthorId(id)
	if notes == nil {
		return httpAdapter.NewErrorResponse(404, "Note not found")
	}
	if err != nil {
		log.Println("[NoteController] Error on get note:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}

	noteVMs := []viewmodels.NoteVM{}
	for _, note := range notes {
		noteVMs = append(noteVMs, viewmodels.NoteVM{}.FromDomain(note))
	}
	return httpAdapter.NewSuccessResponse(200, noteVMs)
}

func GetNoteByCommunityIDController(request httpAdapter.Request) httpAdapter.Response {
	id, ok := request.GetSingleParam("id")
	if !ok {
		log.Println("[NoteController] No id provided on get:")
		return httpAdapter.NewErrorResponse(400, "Invalid community id")
	}

	notes, err := container.NoteService.GetByCommunityId(id)
	if notes == nil {
		return httpAdapter.NewErrorResponse(404, "Note not found")
	}
	if err != nil {
		log.Println("[NoteController] Error on get note:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}

	noteVMs := []viewmodels.NoteVM{}
	for _, note := range notes {
		noteVMs = append(noteVMs, viewmodels.NoteVM{}.FromDomain(note))
	}
	return httpAdapter.NewSuccessResponse(200, noteVMs)
}

func DeleteNoteController(request httpAdapter.Request) httpAdapter.Response {
	id, ok := request.GetSingleParam("id")
	if !ok {
		log.Println("[NoteController] No id provided on delete:")
		return httpAdapter.NewErrorResponse(400, "Invalid note id")
	}

	err := container.NoteService.Delete(id)
	if err != nil {
		log.Println("[NoteController] Error on delete note:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}
	return httpAdapter.NewNoContentResponse()
}

func GetAllNoteController(request httpAdapter.Request) httpAdapter.Response {
	notes, err := container.NoteService.GetAll()
	if err != nil {
		log.Println("[NoteController] Error on get all notes:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}

	return httpAdapter.NewSuccessResponse(http.StatusOK, notes)
}
