package domain

type Note struct {
	Id        		string
	Title     		string
	AuthorID  		string
	Content   		string
	CreatedAt 		string
	UpdatedAt 		string
	LikeCount 		int
	CommentCount  int
}

type FullNote struct {
	Id        	 string
	Title     	 string
	AuthorID  	 string
	Content   	 string
	CreatedAt 	 string
	UpdatedAt 	 string
	LikeCount 	 int
	CommentCount int
	Tags      []NoteTag
	// Communities []domain.Community
}

type FullNoteList struct {
	Id        	 string
	Title     	 string
	AuthorID  	 string
	Author	  	 string
	Content   	 string
	PublishedDate	string
	UpdatedDate 	 string
	LikesCount 	 int
	CommentCount int
	Tags     	 []string
	// Communities []domain.Community
}