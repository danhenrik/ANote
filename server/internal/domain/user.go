package domain

type User struct {
	Id       string
	Email    string
	Password string
}

func (u *User) GetFieldAdresses() []any {
	return []any{&u.Id, &u.Email, &u.Password}
}
