package routes

import (
	httpAdapter "anote/cmd/interfaces"
	"anote/internal/container"
	"anote/internal/viewmodels"
	"encoding/json"
	"log"
	"net/http"
)

func CreateUserController(request httpAdapter.Request) httpAdapter.Response {
	var userVM viewmodels.CreateUserVM
	if err := json.Unmarshal([]byte(request.Body), &userVM); err != nil {
		log.Println("[UserController] Error on create user unmarshal:", err)
		return httpAdapter.NewErrorResponse(400, "Invalid content-type")
	}

	user := userVM.ToDomainUser()
	if err := container.UserService.Create(&user); err != nil {
		log.Println("[UserController] Error on create user:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}
	// move tmp file if it exists to assets folder
	if len(request.Files) != 0 {
		if err := container.UserService.SaveAvatar(user.Id, request.Files[0]); err != nil {
			log.Println("[UserController] Error on save avatar:", err)
			return httpAdapter.NewErrorResponse(201, "Saved user but failed to save avatar")
		}
	}
	return httpAdapter.NewNoContentResponse()
}

func DeleteUserAvatarController(request httpAdapter.Request) httpAdapter.Response {
	if err := container.UserService.DeleteAvatar(request.User.ID); err != nil {
		log.Println("[UserController] Error on delete avatar:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}
	return httpAdapter.NewNoContentResponse()
}

func UpdateUserAvatarController(request httpAdapter.Request) httpAdapter.Response {
	if err := container.UserService.SaveAvatar(request.User.ID, request.Files[0]); err != nil {
		log.Println("[UserController] Error on save avatar:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}
	return httpAdapter.NewNoContentResponse()
}

func GetAllUsersController(request httpAdapter.Request) httpAdapter.Response {
	users, err := container.UserService.GetAll()
	if err != nil {
		log.Println("[UserController] Error on get all users:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}

	var userVMs []viewmodels.UserVM
	for _, user := range users {
		userVMs = append(userVMs, viewmodels.UserVMFromDomainUser(user))
	}
	return httpAdapter.NewSuccessResponse(http.StatusOK, userVMs)
}

func GetUserByUsernameController(request httpAdapter.Request) httpAdapter.Response {
	var userVM any = nil
	userId, ok := request.GetSingleParam("username")
	if ok && userId != "" {
		user, err := container.UserService.GetByUsername(userId)
		if err != nil {
			log.Println("[UserController] Error on get user by username:", err)
			return httpAdapter.NewErrorResponse(err.Status, err.Message)
		}

		if user != nil {
			userVM = viewmodels.UserVMFromDomainUser(*user)
		}
	}
	return httpAdapter.NewSuccessResponse(http.StatusOK, userVM)
}

func GetUserByEmailController(request httpAdapter.Request) httpAdapter.Response {
	var userVM any
	email, ok := request.GetSingleParam("email")
	if ok && email != "" {
		user, err := container.UserService.GetByEmail(email)
		if err != nil {
			log.Println("[UserController] Error on get user by email:", err)
			return httpAdapter.NewErrorResponse(err.Status, err.Message)
		}
		if user != nil {
			userVM = viewmodels.UserVMFromDomainUser(*user)
		}
	}
	return httpAdapter.NewSuccessResponse(http.StatusOK, userVM)
}

func UpdateUserEmailController(request httpAdapter.Request) httpAdapter.Response {
	var userVM viewmodels.UpdateEmailVM
	if err := json.Unmarshal([]byte(request.Body), &userVM); err != nil {
		log.Println("[UserController] Error on update email unmarshal:", err)
		return httpAdapter.NewErrorResponse(400, "Invalid content-type")
	}

	username := request.User.ID
	if err := container.UserService.UpdateEmail(username, userVM.Email); err != nil {
		log.Println("[UserController] Error on update email:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}
	return httpAdapter.NewNoContentResponse()
}

func UpdateUserPasswordController(request httpAdapter.Request) httpAdapter.Response {
	var userVM viewmodels.UpdatePasswordVM
	if err := json.Unmarshal([]byte(request.Body), &userVM); err != nil {
		log.Println("[UserController] Error on update password unmarshal:", err)
		return httpAdapter.NewErrorResponse(400, "Invalid content-type")
	}

	if err := container.UserService.UpdatePassword(
		request.User.ID,
		userVM.OldPassword,
		userVM.NewPassword,
	); err != nil {
		log.Println("[UserController] Error on update password:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}
	return httpAdapter.NewNoContentResponse()
}

func DeleteUserController(request httpAdapter.Request) httpAdapter.Response {
	username := request.User.ID
	if err := container.UserService.Delete(username); err != nil {
		log.Println("[UserController] Error on delete user:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}
	return httpAdapter.NewNoContentResponse()
}

func GetUserAvatarController(request httpAdapter.Request) httpAdapter.Response {
	username, ok := request.GetSingleParam("username")
	if !ok || username == "" {
		return httpAdapter.NewErrorResponse(400, "Invalid username")
	}

	user, err := container.UserService.GetByUsername(username)
	if err != nil {
		log.Println("[UserController] Error on get avatar:", err)
		return httpAdapter.NewErrorResponse(err.Status, err.Message)
	}

	if user.Avatar == nil || *user.Avatar == "" {
		return httpAdapter.NewErrorResponse(404, "Avatar not found")
	}

	var filename string = *user.Avatar
	request.Raw.File("./internal/assets/" + filename)
	return httpAdapter.NewNoContentResponse()
}
