package routes

import (
	httpAdapter "anote/cmd/interfaces"
	"anote/internal/container"
	"anote/internal/domain"
	"anote/internal/viewmodels"
	"encoding/json"
	"log"
)

func CreateCommunityController(request httpAdapter.Request) httpAdapter.Response {
	var communityVM viewmodels.CreateCommunityVM
	if err := json.Unmarshal([]byte(request.Body), &communityVM); err != nil {
		log.Println("[CommunityController] Error on create tag unmarshal:", err)
		return httpAdapter.NewErrorResponse(400, "Invlaid content-type")
	}
	community := domain.Community{Name: communityVM.Name}
	if err := container.CommunityService.Create(&community); err != nil {
		log.Println("[CommunityController] Error on create tag:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}
	return httpAdapter.NewSuccessResponse(201, map[string]string{"id": community.Id})
}

func GetAllCommunitiesController(request httpAdapter.Request) httpAdapter.Response {
	communities, err := container.CommunityService.GetAll()
	if err != nil {
		log.Println("[CommunityController] Error on get all communities:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}
	return httpAdapter.NewSuccessResponse(200, communities)
}

func DeleteCommunityController(request httpAdapter.Request) httpAdapter.Response {
	id, ok := request.GetSingleParam("id")
	if !ok {
		log.Println("[CommunityController] Error on delete tag: id not found")
		return httpAdapter.NewErrorResponse(400, "id not found")
	}

	if err := container.CommunityService.Delete(id); err != nil {
		log.Println("[CommunityController] Error on delete tag:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}
	return httpAdapter.NewNoContentRespone()
}

func JoinCommunityController(request httpAdapter.Request) httpAdapter.Response {
	id, ok := request.GetSingleParam("id")
	if !ok {
		log.Println("[CommunityController] Error on join community: id not found")
		return httpAdapter.NewErrorResponse(400, "id not found")
	}

	err := container.CommunityService.Join(id, request.User.ID)
	if err != nil {
		log.Println("[CommunityController] Error on join community: user not found")
		return httpAdapter.NewErrorResponse(500, "Could not join community")
	}
	return httpAdapter.NewNoContentRespone()
}

func LeaveCommunityController(request httpAdapter.Request) httpAdapter.Response {
	id, ok := request.GetSingleParam("id")
	if !ok {
		log.Println("[CommunityController] Error on leave community: id not found")
		return httpAdapter.NewErrorResponse(400, "id not found")
	}

	err := container.CommunityService.Leave(id, request.User.ID)
	if err != nil {
		log.Println("[CommunityController] Error on leave community: user not found")
		return httpAdapter.NewErrorResponse(500, "Could not leave community")
	}
	return httpAdapter.NewNoContentRespone()
}
