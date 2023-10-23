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
	tag := domain.Tag{Name: tagVM.Name}

	if err := container.NoteTagService.Create(&tag); err != nil {
		log.Println("[NoteTagController] Error on create tag:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}
	return httpAdapter.NewSuccessResponse(201, map[string]string{"id": tag.Id})
}

func GetAllTagsController(request httpAdapter.Request) httpAdapter.Response {
	tags, err := container.NoteTagService.GetAll()
	if err != nil {
		log.Println("[NoteTagController] Error on get tags:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}
	return httpAdapter.NewSuccessResponse(200, tags)
}

func DeleteTagController(request httpAdapter.Request) httpAdapter.Response {
	id, ok := request.GetSingleParam("id")
	if !ok {
		log.Println("[NoteTagController] Error on delete tag: id not found")
		return httpAdapter.NewErrorResponse(400, "id not found")
	}

	if err := container.NoteTagService.Delete(id); err != nil {
		log.Println("[NoteTagController] Error on delete tag:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}
	return httpAdapter.NewNoContentRespone()
}
