package routes

import (
	httpAdapter "anote/cmd/interfaces"
	"anote/internal/container"
	"anote/internal/viewmodels"
	"encoding/json"
	"log"
	"net/http"
)

func CreateCommentController(request httpAdapter.Request) httpAdapter.Response {
	var commentVM viewmodels.CreateCommentVM
	if err := json.Unmarshal([]byte(request.Body), &commentVM); err != nil {
		log.Println("[CommentController] Error on create comment unmarshal:", err)
		return httpAdapter.NewErrorResponse(400, "Invalid content-type")
	}

	comment := commentVM.ToDomainComment()
	if err := container.CommentService.Create(&comment); err != nil {
		log.Println("[CommentController] Error on create comment:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}
	return httpAdapter.NewNoContentRespone()
}

func DeleteCommentController(request httpAdapter.Request) httpAdapter.Response {
	idComment, okIdComment := request.GetSingleParam("id")

	if !okIdComment {
		log.Println("[CommentController] Error on delete comment: id not found")
		return httpAdapter.NewErrorResponse(400, "id not found")
	}
	if err := container.CommentService.Delete(idComment); err != nil {
		log.Println("[CommentController] Error on delete comment:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}
	return httpAdapter.NewNoContentRespone()
}

func GetNoteCommentsController(request httpAdapter.Request) httpAdapter.Response {
	idNote, okIdNote := request.GetSingleParam("idNote")

	if !okIdNote {
		log.Println("[CommentController] Error on get comment: id not found")
		return httpAdapter.NewErrorResponse(400, "id not found")
	}

	comments, err := container.CommentService.GetNoteComments(idNote)

	if err != nil {
		log.Println("[CommentController] Error on get note comments:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}

	return httpAdapter.NewSuccessResponse(http.StatusOK, comments)
}