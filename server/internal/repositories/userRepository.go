package repositories

import (
	"anote/internal/domain"
	interfaces "anote/internal/ports/database"
	"errors"
	"fmt"
	"log"
)

// this component is carried with making the SQL queries to the database
type UserRepository struct{ DBConn interfaces.DBConnection }

func NewUserRepository(DBConn interfaces.DBConnection) UserRepository {
	return UserRepository{
		DBConn: DBConn,
	}
}

func (this UserRepository) Create(user domain.User) error {
	err := this.DBConn.Exec("INSERT INTO users (id, email, password) VALUES ($1, $2, $3)", user.Id, user.Email, user.Password)
	if err != nil {
		log.Println("Error on insert new user: ", err)
		return errors.New(fmt.Sprintf("Error on insert new user: %v", err))
	}
	return nil
}

func (this UserRepository) GetByUsername(username string) (domain.User, error) {
	var user domain.User
	err := this.DBConn.QueryOne(&user, "SELECT * FROM users WHERE id = $1", username)
	if err != nil {
		errMessage := fmt.Sprintf("Error on get user by username: %v", err)
		log.Println(errMessage)
		return user, errors.New(errMessage)
	}

	return user, nil
}

func (this UserRepository) GetByEmail(email string) (domain.User, error) {
	var user domain.User
	err := this.DBConn.QueryOne(&user, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		errMessage := fmt.Sprintf("Error on get user by email: %v", err)
		log.Println(errMessage)
		return user, errors.New(errMessage)
	}

	return user, nil
}

func (this UserRepository) GetAll() ([]domain.User, error) {
	var users []domain.User
	// err := this.DBConn.QueryMultiple(users, "SELECT * FROM users")
	// if err != nil {
	// 	errMessage := fmt.Sprintf("Error on get all users: %v", err)
	// 	log.Println(errMessage)
	// 	return users, errors.New(errMessage)
	// }

	return users, nil
}

func (_ UserRepository) Update(user domain.User) error {
	return nil
}

func (this UserRepository) Delete(username string) error {
	return this.DBConn.Exec("DELETE FROM users WHERE id = $1", username)
}
