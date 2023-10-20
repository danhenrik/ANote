package viewmodels

import "anote/internal/domain"

type CreateUserVM struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateEmailVM struct {
	Email string `json:"email"`
}

type UpdatePasswordVM struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type ChangePasswordVM struct {
	Token       string `json:"token"`
	NewPassword string `json:"new_password"`
}

type UserVM struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (this *CreateUserVM) ToDomainUser() domain.User {
	var user domain.User
	user.Id = this.Username
	user.Email = this.Email
	user.Password = this.Password
	return user
}

func UserVMFromDomainUser(userD domain.User) UserVM {
	var user UserVM
	user.Username = userD.Id
	user.Email = userD.Email
	return user
}

func (this *UserVM) ToDomainUser() domain.User {
	var user domain.User
	user.Id = this.Username
	user.Email = this.Email
	return user
}
