package domain

type Note struct {
	Id        string
	Title     string
	AuthorID  string
	Content   string
	CreatedAt string
	UpdatedAt string
}

type FullNote struct {
	Id          string
	Title       string
	AuthorID    string
	Content     string
	CreatedAt   string
	UpdatedAt   string
	Tags        []Tag
	Communities []Community
}

type Community struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Tag struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Like struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
	NoteId string `json:"note_id"`
	CreatedAt string `json:"created_at"`
}

type Comment struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	NoteId    string `json:"note_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

type NoteComment struct {
	Id        string `json:"Id"`
	Author    string `json:"Author"`
	Content   string `json:"Content"`
	CreatedAt string `json:"CreatedAt"`
}

type FilteredNote struct {
	Id            string      `json:"id"`
	Title         string      `json:"title"`
	Content       string      `json:"content"`
	PublishedDate string      `json:"published_date"`
	UpdatedDate   string      `json:"updated_date"`
	Author        string      `json:"author"`
	Communities   []Community `json:"communities"`
	Tags          []Tag       `json:"tags"`
	LikesCount    int         `json:"likes_count"`
	Likes         []Like      `json:"likes"`
	CommentCount  int         `json:"comment_count"`
	Commenters    []Comment   `json:"commenters"`
}
