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

	disclaimer := ""
	if len(request.Files) != 0 {
		if err := container.CommunityService.SaveBackground(community.Id, request.Files[0]); err != nil {
			log.Println("[CommunityController] Error on save community background:", err)
			disclaimer = "Saved community but failed to save background"
		}
	}

	err := container.CommunityService.Join(community.Id, request.User.ID)
	if err != nil {
		log.Println("[CommunityController] Error on join community")
		if disclaimer != "" {
			disclaimer = disclaimer + "; Also could not join community"
		} else {
			disclaimer = "Could not join community"
		}
	}

	response := map[string]string{"id": community.Id}
	if disclaimer != "" {
		response["disclaimer"] = disclaimer
	}
	return httpAdapter.NewSuccessResponse(201, response)
}

func DeleteCommunityBackgroundController(request httpAdapter.Request) httpAdapter.Response {
	id, ok := request.GetSingleParam("id")
	if !ok {
		log.Println("[CommunityController] Error on delete background: community id not found")
		return httpAdapter.NewErrorResponse(400, "community id not found")
	}

	if err := container.CommunityService.DeleteBackground(id); err != nil {
		log.Println("[CommunityController] Error on delete community background:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}
	return httpAdapter.NewNoContentRespone()
}

func UpdateCommunityBackgroundController(request httpAdapter.Request) httpAdapter.Response {
	id, ok := request.GetSingleParam("id")
	if !ok {
		log.Println("[CommunityController] Error on delete background: community id not found")
		return httpAdapter.NewErrorResponse(400, "community id not found")
	}

	if err := container.CommunityService.DeleteBackground(id); err != nil {
		log.Println("[CommunityController] Error on delete to update community background:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}

	if err := container.CommunityService.SaveBackground(id, request.Files[0]); err != nil {
		log.Println("[CommunityController] Error on save community background:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}
	return httpAdapter.NewNoContentRespone()
}

func GetAllCommunitiesController(request httpAdapter.Request) httpAdapter.Response {
	communities, err := container.CommunityService.GetAll()
	if err != nil {
		log.Println("[CommunityController] Error on get all communities:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}
	return httpAdapter.NewSuccessResponse(200, communities)
}

func GetCurrentUserCommunities(request httpAdapter.Request) httpAdapter.Response {
	userId := request.User.ID
	communities, err := container.CommunityService.GetByUserId(userId)
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
