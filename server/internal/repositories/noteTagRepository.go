package repositories

import (
	"anote/internal/domain"
	"anote/internal/errors"
	"anote/internal/interfaces"
	"log"
	"reflect"
)

type NoteTagRepository struct {
	DBConn interfaces.DBConnection
}

func NewNoteTagRepository(
	DBConn interfaces.DBConnection,
) NoteTagRepository {
	return NoteTagRepository{DBConn: DBConn}
}

func (this NoteTagRepository) Create(tag *domain.NoteTag) *errors.AppError {
	err := this.DBConn.Exec(
		"INSERT INTO tags (id, name) VALUES ($1, $2)",
		tag.Id,
		tag.Name,
	)
	if err != nil {
		log.Println("[NoteTagRepo] Error on insert new tag:", err)
		return err
	}
	return nil
}

func (this NoteTagRepository) GetById(id string) (*domain.NoteTag, *errors.AppError) {
	objType := reflect.TypeOf(domain.NoteTag{})
	res, err := this.DBConn.QueryOne(objType, "SELECT * FROM tags WHERE id = $1", id)
	if err != nil {
		log.Println("[NoteTagRepo] Error on get tag by id:", err)
		return nil, err
	}

	if tag, ok := res.(domain.NoteTag); ok {
		return &tag, nil
	}
	return nil, nil
}
