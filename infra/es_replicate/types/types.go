package types

// Shared
type Update struct {
	WalId  []uint8
	Change []struct {
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
	} `json:"change"`
}

type Community struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Tag struct {
	Id    string `json:"id"`
	TagId string `json:"tag_id"`
	Name  string `json:"name"`
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
	Author        string      `json:"author"` // username (id)
	Communities   []Community `json:"communities"`
	Tags          []Tag       `json:"tags"`
	LikesCount    int         `json:"likes_count"`
	Likes         []Like      `json:"likes"` // usernames (id)
	CommentCount  int         `json:"comment_count"`
	Commenters    []Comment   `json:"commenters"` // usernames (id)
}
