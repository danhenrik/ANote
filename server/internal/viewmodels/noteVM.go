package viewmodels

import (
	"anote/internal/domain"
)

type CreateNoteVM struct {
	Title       string   `json:"title"`
	Content     string   `json:"content"`
	Tags        []string `json:"tags"`
	Communities []string `json:"communities"`
}

type UpdateNoteVM struct {
	Id                string   `json:"id"`
	Title             string   `json:"title"`
	Content           string   `json:"content"`
	AddTags           []string `json:"add_tags"`
	RemoveTags        []string `json:"remove_tags"`
	AddCommunities    []string `json:"add_communities"`
	RemoveCommunities []string `json:"remove_communities"`
}

type NoteVM struct {
	Id           string        `json:"id"`
	Title        string        `json:"title"`
	Content      string        `json:"content"`
	AuthorID     string        `json:"author_id"`
	CreatedAt    string        `json:"created_at"`
	UpdatedAt    string        `json:"updated_at"`
	Tags         []NoteTagVM   `json:"tags"`
	Communities  []CommunityVM `json:"communities"`
	Comments     []CommentVM   `json:"comments"`
	Likes        []LikeVM      `json:"likes"`
	LikeCount    int           `json:"like_count"`
	CommentCount int           `json:"comment_count"`
}

func (note CreateNoteVM) ToDomainNote() domain.Note {
	return domain.Note{
		Title:   note.Title,
		Content: note.Content,
	}
}

func (n NoteVM) FromDomain(note domain.FullNote) NoteVM {
	n = NoteVM{
		Id:          note.Id,
		Title:       note.Title,
		Content:     note.Content,
		AuthorID:    note.AuthorID,
		CreatedAt:   note.CreatedAt,
		UpdatedAt:   note.UpdatedAt,
		Tags:        []NoteTagVM{},
		Communities: []CommunityVM{},
		Comments:    []CommentVM{},
		Likes:       []LikeVM{},
	}
	for _, tag := range note.Tags {
		n.Tags = append(n.Tags, NoteTagVM{
			Id:   tag.Id,
			Name: tag.Name,
		})
	}
	for _, community := range note.Communities {
		n.Communities = append(n.Communities, CommunityVM{
			Id:   community.Id,
			Name: community.Name,
		})
	}
	for _, comment := range note.Comments {
		n.Comments = append(n.Comments, CommentVM{
			UserId:    comment.UserId,
			NoteId:    comment.NoteId,
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt,
		})
	}
	for _, like := range note.Likes {
		n.Likes = append(n.Likes, LikeVM{
			UserId:    like.UserId,
			NoteId:    like.NoteId,
			CreatedAt: like.CreatedAt,
		})
	}
	n.LikeCount = len(n.Likes)
	n.CommentCount = len(n.Comments)

	return n
}
