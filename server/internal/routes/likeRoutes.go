package routes

import (
	httpAdapter "anote/cmd/interfaces"
	"anote/internal/container"
	"anote/internal/viewmodels"
	"encoding/json"
	"log"
	"net/http"
)

func CreateLikeController(request httpAdapter.Request) httpAdapter.Response {
	var likeVM viewmodels.CreateLikeVM
	if err := json.Unmarshal([]byte(request.Body), &likeVM); err != nil {
		log.Println("[LikeController] Error on create like unmarshal:", err)
		return httpAdapter.NewErrorResponse(400, "Invalid content-type")
	}

	like := likeVM.ToDomainLike()
	if err := container.LikeService.Create(&like); err != nil {
		log.Println("[LikeController] Error on create like:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}
	return httpAdapter.NewNoContentRespone()
}

func DeleteLikeController(request httpAdapter.Request) httpAdapter.Response {
	idUser, okIdUser := request.GetSingleParam("idUser")
	idNote, okIdNote := request.GetSingleParam("idNote")

	if !okIdUser || !okIdNote {
		log.Println("[LikeController] Error on delete like: id not found")
		return httpAdapter.NewErrorResponse(400, "id not found")
	}
	if err := container.LikeService.Delete(idUser, idNote); err != nil {
		log.Println("[LikeController] Error on delete like:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}
	return httpAdapter.NewNoContentRespone()
}

func GetLikeController(request httpAdapter.Request) httpAdapter.Response {
	idUser, okIdUser := request.GetSingleParam("idUser")
	idNote, okIdNote := request.GetSingleParam("idNote")

	if !okIdUser || !okIdNote {
		log.Println("[LikeController] Error on get like: id not found")
		return httpAdapter.NewErrorResponse(400, "id not found")
	}

	like, err := container.LikeService.GetByIdUserAndIdNote(idUser, idNote)

	if err != nil {
		log.Println("[LikeController] Error on get like:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}

	return httpAdapter.NewSuccessResponse(http.StatusOK, like)

}

func CountLikeByIdNoteController(request httpAdapter.Request) httpAdapter.Response {
	idNote, okIdNote := request.GetSingleParam("idNote")

	if !okIdNote {
		log.Println("[LikeController] Error on count like: id not found")
		return httpAdapter.NewErrorResponse(400, "id not found")
	}

	numberLikes, err := container.LikeService.CountLikeByIdNoteController(idNote)

	if err != nil {
		log.Println("[LikeController] Error on count like:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}

	return httpAdapter.NewSuccessResponse(http.StatusOK, numberLikes)
}