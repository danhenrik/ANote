package viewmodels

import (
	"anote/internal/domain"
)

type CreateNoteVM struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
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
	Id        string      `json:"id"`
	Title     string      `json:"title"`
	Content   string      `json:"content"`
	AuthorID  string      `json:"author_id"`
	CreatedAt string      `json:"created_at"`
	UpdatedAt string      `json:"updated_at"`
	LikeCount 	 int      `json:"LikesCount"`
	CommentCount int      `json:"CommentCount"`
	Tags      []NoteTagVM `json:"tags"`
}

type NoteListVM struct {
	Id        	 string   `json:"Id"`
	Title     	 string   `json:"Title"`
	Content   	 string   `json:"Content"`
	AuthorID  	 string   `json:"author_id"`
	Author			 string		`json:"Author"`
	PublishedDate 	 string   `json:"PublishedDate"`
	UpdatedDate 	 string   `json:"UpdatedDate"`
	LikesCount 	 int      `json:"LikesCount"`
	CommentCount int      `json:"CommentCount"`
	Tags      	 []string	`json:"Tags"`
}

func (note CreateNoteVM) ToDomainNote() domain.Note {
	return domain.Note{
		Title:   note.Title,
		Content: note.Content,
	}
}

func (n NoteVM) FromDomain(note domain.FullNote) NoteVM {
	n = NoteVM{
		Id:        		note.Id,
		Title:     		note.Title,
		Content:   		note.Content,
		AuthorID:  		note.AuthorID,
		CreatedAt: 		note.CreatedAt,
		UpdatedAt: 		note.UpdatedAt,
		LikeCount: 		note.LikeCount,
		CommentCount: note.CommentCount,
		Tags:      		[]NoteTagVM{},
	}	
	for _, tag := range note.Tags {
		n.Tags = append(n.Tags, NoteTagVM{
			Id:   tag.Id,
			Name: tag.Name,
		})
	}
	return n
}
