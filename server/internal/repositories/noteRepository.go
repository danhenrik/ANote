package repositories

import (
	"anote/internal/domain"
	"anote/internal/errors"
	"anote/internal/interfaces"
	"log"
)

type NoteRepository struct{ DBConn interfaces.DBConnection }

func NewNoteRepository(DBConn interfaces.DBConnection) NoteRepository {
	return NoteRepository{
		DBConn: DBConn,
	}
}

func (this NoteRepository) Create(note *domain.Note) *errors.AppError {
	err := this.DBConn.Exec(
		"INSERT INTO notes (id, title, author_id, content) VALUES ($1, $2, $3, $4)",
		note.Id,
		note.Title,
		note.AuthorID,
		note.Content,
	)
	if err != nil {
		log.Println("[NoteRepo] Error on insert new note:", err)
		return err
	}
	return nil
}

func (this NoteRepository) GetByID(id string) (*domain.Note, *errors.AppError) {
	return nil, nil
}

func (this NoteRepository) GetByTitle(title string) ([]domain.Note, *errors.AppError) {
	return nil, nil

}

func (this NoteRepository) Update(user *domain.Note) *errors.AppError {
	return nil

}

func (this NoteRepository) Delete(username string) *errors.AppError {
	return nil
}
