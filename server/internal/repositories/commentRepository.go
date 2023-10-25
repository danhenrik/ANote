package repositories

import (
	"anote/internal/domain"
	"anote/internal/errors"
	"anote/internal/interfaces"
	"log"
	"reflect"
)

type CommentRepository struct {
	DBConn interfaces.DBConnection
}

func NewCommentRepository(
	DBConn interfaces.DBConnection,
) CommentRepository {
	return CommentRepository{DBConn: DBConn}
}

func (this CommentRepository) Create(comment *domain.Comment) *errors.AppError {
	err := this.DBConn.Exec(
		"INSERT INTO comments (id, user_id, note_id, content) VALUES ($1, $2, $3, $4)",
		comment.Id,
		comment.UserId,
		comment.NoteId,
		comment.Content,
	)
	if err != nil {
		log.Println("[CommentRepo] Error on insert new comment:", err)
		return err
	}
	return nil
}

func (this CommentRepository) Delete(idComment string) *errors.AppError {
	err := this.DBConn.Exec("DELETE FROM comments WHERE id = $1", idComment)
	if err != nil {
		log.Println("[CommentRepo] Error on delete comment:", err)
		return err
	}
	return nil
}

func (this CommentRepository) GetNoteComments(idNote string) ([]domain.Comment, *errors.AppError) {
	objType := reflect.TypeOf(domain.Comment{})

	res, err := this.DBConn.QueryMultiple(objType, "SELECT * FROM comments WHERE note_id = $1", idNote)

	if err != nil {
		log.Println("[CommentRepo] Error on get note comments:", err)
		return []domain.Comment{}, err
	}

	if comments, ok := res.([]domain.Comment); ok {
		return comments, nil
	}

	return []domain.Comment{}, nil
}