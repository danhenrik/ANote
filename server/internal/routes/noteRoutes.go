package routes

import (
	httpAdapter "anote/cmd/interfaces"
	"anote/internal/container"
	"anote/internal/viewmodels"
	"encoding/json"
	"log"
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
