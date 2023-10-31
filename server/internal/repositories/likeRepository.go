package repositories

import (
	"anote/internal/domain"
	"anote/internal/errors"
	"anote/internal/helpers"
	"anote/internal/interfaces"
	"log"
	"reflect"
)

type LikeRepository struct {
	DBConn interfaces.DBConnection
}

func NewLikeRepository(
	DBConn interfaces.DBConnection,
) LikeRepository {
	return LikeRepository{DBConn: DBConn}
}

func (this LikeRepository) Create(like *domain.Like) *errors.AppError {
	err := this.DBConn.Exec(
		"INSERT INTO likes (id, user_id, note_id) VALUES ($1, $2, $3)",
		helpers.NewUUID(),
		like.UserId,
		like.NoteId,
	)
	if err != nil {
		log.Println("[LikeRepo] Error on insert new like:", err)
		return err
	}
	return nil
}

func (this LikeRepository) Delete(idUser string, idNote string) *errors.AppError {
	err := this.DBConn.Exec("DELETE FROM likes WHERE user_id = $1 AND note_id = $2", idUser, idNote)
	if err != nil {
		log.Println("[LikeRepo] Error on delete like:", err)
		return err
	}
	return nil
}

func (this LikeRepository) GetByIdUserAndIdNote(idUser string, idNote string) (*domain.Like, *errors.AppError) {
	objType := reflect.TypeOf(domain.Like{})

	res, err := this.DBConn.QueryOne(objType, "SELECT * FROM likes WHERE user_id = $1 AND note_id = $2", idUser, idNote)

	if err != nil {
		log.Println("[LikeRepo] Error on get like:", err)
		return nil, err
	}
	if res == nil {
		return nil, nil
	}

	if like, ok := res.(domain.Like); ok {
		return &like, nil
	}
	log.Println("[LikeRepo] Like not found")
	return nil, nil

}
func (this LikeRepository) GetByIdNote(idNote string) ([]domain.Like, *errors.AppError) {
	objType := reflect.TypeOf(domain.Like{})

	res, err := this.DBConn.QueryMultiple(objType, "SELECT * FROM likes WHERE note_id = $1", idNote)

	if err != nil {
		log.Println("[LikeRepo] Error on get like:", err)
		return []domain.Like{}, err
	}
	if res == nil {
		return []domain.Like{}, nil
	}

	if likes, ok := res.([]domain.Like); ok {
		return likes, nil
	}
	log.Println("[LikeRepo] Like not found")
	return []domain.Like{}, nil

}
