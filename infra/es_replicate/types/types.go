package types

// Shared
type Change struct {
	Kind         string   `json:"kind"`
	Schema       string   `json:"schema"`
	Table        string   `json:"table"`
	ColumnNames  []string `json:"columnnames"`
	ColumnTypes  []string `json:"columntypes"`
	ColumnValues []string `json:"columnvalues"`
	OldKeys      struct {
		KeyNames  []string `json:"keynames"`
		KeyTypes  []string `json:"keytypes"`
		KeyValues []string `json:"keyvalues"`
	} `json:"oldkeys"`
}

type Update struct {
	WalId  []uint8
	Change []Change `json:"change"`
}

type Community struct {
	RId  string `json:"r_id"`
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Tag struct {
	RId  string `json:"r_id"`
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Like struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
}

type Comment struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
}

type Note struct {
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
