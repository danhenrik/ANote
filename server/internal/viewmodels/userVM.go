package viewmodels

import "anote/internal/domain"

type UserVM struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password,omitempty" form:"password"`
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
	user.Password = this.Password
	return user
}

func (this *UserVM) IsEmpty() bool {
	return this.Username == "" || this.Email == "" || this.Password == ""
}
