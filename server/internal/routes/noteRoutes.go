package routes

import (
	httpAdapter "anote/cmd/interfaces"
	"anote/internal/container"
	"anote/internal/viewmodels"
	"encoding/json"
	"log"
	"net/http"
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
	return httpAdapter.NewNoContentRespone()
}

func SearchNoteController(request httpAdapter.Request) httpAdapter.Response {
	title, _ := request.GetSingleQuery("title")
	content, _ := request.GetSingleQuery("content")
	tags, _ := request.GetQuerySlice("tags")
	author, _ := request.GetSingleQuery("author")
	communities, _ := request.GetQuerySlice("communities")
	published_date, _ := request.GetSingleQuery("date")
	from_published_date, _ := request.GetSingleQuery("from_date")
	to_published_date, _ := request.GetSingleQuery("to_date")

	createdAt := [2]string{"", ""}
	if published_date != "" {
		createdAt[0] = published_date
	} else if from_published_date != "" && to_published_date != "" {
		createdAt[0] = from_published_date
		createdAt[1] = to_published_date
	}

	notes, err := container.NoteService.Search(
		title,
		content,
		author,
		tags,
		communities,
		createdAt,
	)
	if err != nil {
		log.Println("[NoteController] Error on search note:", err)
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

	// check access
	noteVM := viewmodels.NoteVM{}.FromDomain(*note)
	return httpAdapter.NewSuccessResponse(200, noteVM)
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
	return httpAdapter.NewNoContentRespone()
}

func GetAllNoteController(request httpAdapter.Request) httpAdapter.Response {
	notes, err := container.NoteService.GetAll()
	if err != nil {
		log.Println("[NoteController] Error on get all notes:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}

	return httpAdapter.NewSuccessResponse(http.StatusOK, notes)
}
