package domain

type User struct {
	Id        string
	Email     string
	Password  *string
	Google_id *string
	CreatedAt string
	Avatar    *string
}
