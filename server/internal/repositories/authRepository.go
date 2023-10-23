package repositories

import (
	"anote/internal/domain"
	"anote/internal/errors"
	"anote/internal/interfaces"
	"log"
	"reflect"
)

type AuthRepository struct{ DBConn interfaces.DBConnection }

func NewAuthRepository(DBConn interfaces.DBConnection) AuthRepository {
	return AuthRepository{
		DBConn: DBConn,
	}
}

func (this AuthRepository) SaveToken(token string, userId string) *errors.AppError {
	// should actually use redis or something like that for expiration
	err := this.DBConn.Exec(
		"INSERT INTO tokens (token, user_id) VALUES ($1, $2)",
		token,
		userId,
	)
	if err != nil {
		log.Println("[AuthRepo] Error on save new token:", err)
		return err
	}
	return nil
}

type Token struct {
	Id     string `json:"id"`
	Token  string `json:"token"`
	UserId string `json:"user_id"`
}

func (this AuthRepository) RetrieveToken(token string) (*domain.User, *errors.AppError) {
	objType := reflect.TypeOf(Token{})
	tokenRes, err := this.DBConn.QueryOne(objType, "SELECT * FROM tokens WHERE token = $1", token)
	if err != nil {
		log.Println("[AuthRepo] Error retrieving token:", err)
		return nil, err
	}
	if tokenRes == nil {
		return nil, errors.NewAppError(404, "Token not found")
	}
	t := tokenRes.(Token)
	userFromDB, err := this.DBConn.QueryOne(reflect.TypeOf(domain.User{}), "SELECT * FROM users WHERE id = $1", t.UserId)
	if err != nil {
		log.Println("[AuthRepo] Error retrieving user referenced by token:", err)
		return nil, err
	}
	if userFromDB == nil {
		return nil, errors.NewAppError(400, "User related to token not found")
	}
	user := userFromDB.(domain.User)
	return &user, nil
}

func (this AuthRepository) DeleteToken(token string) *errors.AppError {
	err := this.DBConn.Exec("DELETE FROM tokens WHERE token = $1", token)
	if err != nil {
		log.Println("[AuthRepo] Error on delete token:", err)
		return err
	}
	return nil
}
