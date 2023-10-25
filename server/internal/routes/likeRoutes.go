package routes

import (
	httpAdapter "anote/cmd/interfaces"
	"anote/internal/container"
	"anote/internal/viewmodels"
	"encoding/json"
	"log"
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