package routes

import (
	httpAdapter "anote/cmd/interfaces"
	"anote/internal/container"
	"anote/internal/domain"
	"anote/internal/viewmodels"
	"encoding/json"
	"log"
)

func CreateTagController(request httpAdapter.Request) httpAdapter.Response {
	var tagVM viewmodels.CreateNoteTagVM
	if err := json.Unmarshal([]byte(request.Body), &tagVM); err != nil {
		log.Println("[NoteTagController] Error on create tag unmarshal:", err)
		return httpAdapter.NewErrorResponse(400, "Invlaid content-type")
	}
	tag := domain.NoteTag{Name: tagVM.Name}

	if err := container.NoteTagService.Create(&tag); err != nil {
		log.Println("[NoteTagController] Error on create tag:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}
	return httpAdapter.NewSuccessResponse(201, map[string]string{"id": tag.Id})
}
