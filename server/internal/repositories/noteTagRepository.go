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

func (this NoteTagRepository) Create(tag *domain.Tag) *errors.AppError {
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

func (this NoteTagRepository) GetAll() ([]domain.Tag, *errors.AppError) {
	objType := reflect.TypeOf(domain.Tag{})
	res, err := this.DBConn.QueryMultiple(objType, "SELECT * FROM tags")
	if err != nil {
		log.Println("[NoteTagRepo] Error on get all tags:", err)
		return nil, err
	}

	if tags, ok := res.([]domain.Tag); ok {
		return tags, nil
	}
	return []domain.Tag{}, nil
}

func (this NoteTagRepository) GetById(id string) (*domain.Tag, *errors.AppError) {
	objType := reflect.TypeOf(domain.Tag{})
	res, err := this.DBConn.QueryOne(objType, "SELECT * FROM tags WHERE id = $1", id)
	if err != nil {
		log.Println("[NoteTagRepo] Error on get tag by id:", err)
		return nil, err
	}
	if res == nil {
		return nil, nil
	}

	if tag, ok := res.(domain.Tag); ok {
		return &tag, nil
	}
	return nil, nil
}

func (this NoteTagRepository) GetByName(tagName string) (*domain.Tag, *errors.AppError) {
	objType := reflect.TypeOf(domain.Tag{})
	res, err := this.DBConn.QueryOne(objType, "SELECT * FROM tags WHERE name = $1", tagName)
	if err != nil {
		log.Println("[NoteTagRepo] Error on get tag by name:", err)
		return nil, err
	}
	if res == nil {
		return nil, nil
	}

	if tag, ok := res.(domain.Tag); ok {
		return &tag, nil
	}
	return nil, nil
}

func (this NoteTagRepository) GetByNoteId(noteId string) ([]domain.Tag, *errors.AppError) {
	objType := reflect.TypeOf(domain.Tag{})

	res, err := this.DBConn.QueryMultiple(
		objType,
		`SELECT tags.* FROM tags
		INNER JOIN note_tags ON note_tags.tag_id = tags.id
		WHERE note_tags.note_id = $1`,
		noteId,
	)

	if err != nil {
		log.Println("[NoteTagRepo] Error on get tags by note id:", err)
		return nil, err
	}

	if tags, ok := res.([]domain.Tag); ok {
		return tags, nil
	}
	return []domain.Tag{}, nil
}

func (this NoteTagRepository) Delete(id string) *errors.AppError {
	err := this.DBConn.Exec("DELETE FROM tags WHERE id = $1", id)
	if err != nil {
		log.Println("[NoteTagRepo] Error on delete tag:", err)
		return err
	}
	return nil
}
