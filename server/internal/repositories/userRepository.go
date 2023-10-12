package repositories

import (
	"anote/internal/domain"
	"anote/internal/errors"
	"anote/internal/interfaces"
	"log"
	"reflect"
)

// this component is carried with making the SQL queries to the database
type UserRepository struct{ DBConn interfaces.DBConnection }

func NewUserRepository(DBConn interfaces.DBConnection) UserRepository {
	return UserRepository{
		DBConn: DBConn,
	}
}

func (this UserRepository) Create(user *domain.User) *errors.AppError {
	err := this.DBConn.Exec("INSERT INTO users (id, email, password) VALUES ($1, $2, $3)", user.Id, user.Email, user.Password)
	if err != nil {
		log.Println("[UserRepo] Error on insert new user:", err)
		return err
	}
	return nil
}

func (this UserRepository) GetByUsername(username string) (*domain.User, *errors.AppError) {
	objType := reflect.TypeOf(domain.User{})

	res, err := this.DBConn.QueryOne(objType, "SELECT * FROM users WHERE id = $1", username)
	if err != nil {
		log.Println("[UserRepo] Error on get user by username:", err)
		return nil, err
	}

	if user, ok := res.(domain.User); ok {
		return &user, nil
	}
	return nil, nil
}

func (this UserRepository) GetByEmail(email string) (*domain.User, *errors.AppError) {
	objType := reflect.TypeOf(domain.User{})

	res, err := this.DBConn.QueryOne(objType, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		log.Println("[UserRepo] Error on get user by email:", err)
		return nil, err
	}

	if user, ok := res.(domain.User); ok {
		return &user, nil
	}
	return nil, nil
}

func (this UserRepository) GetAll() ([]domain.User, *errors.AppError) {
	objType := reflect.TypeOf(domain.User{})

	res, err := this.DBConn.QueryMultiple(objType, "SELECT * FROM users")
	if err != nil {
		log.Println("[UserRepo] Error on get all users:", err)
		return []domain.User{}, err
	}

	if users, ok := res.([]domain.User); ok {
		return users, nil
	}
	return []domain.User{}, nil
}

func (_ UserRepository) Update(user *domain.User) *errors.AppError {
	return nil
}

func (this UserRepository) Delete(username string) *errors.AppError {
	this.DBConn.Exec("DELETE FROM users WHERE id = $1", username)
	return nil
}

func (this UserRepository) GetUserWithPassword(key string) (*domain.User, *errors.AppError) {
	objType := reflect.TypeOf(domain.User{})

	res, err := this.DBConn.QueryOne(objType, "SELECT * FROM users WHERE id = $1 OR email = $1", key)
	if err != nil {
		log.Println("[UserRepo] Error on get user with password:", err)
		return nil, err
	}

	if user, ok := res.(domain.User); ok {
		return &user, nil
	}
	return nil, nil
}
