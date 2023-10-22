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

	noteId, err := container.NoteService.Create(&note, noteVM.Tags)
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

func GetNoteController(request httpAdapter.Request) httpAdapter.Response {
	return httpAdapter.NewSuccessResponse(200, "Hello from get note")
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