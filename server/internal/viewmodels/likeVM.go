package viewmodels

import "anote/internal/domain"

type CreateLikeVM struct {
	UserId string `json:"user_id"`
	NoteId string `json:"note_id"`
}

func (this *CreateLikeVM) ToDomainLike() domain.Like {
	var like domain.Like
	like.UserId = this.UserId
	like.NoteId = this.NoteId
	return like
}