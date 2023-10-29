package viewmodels

type CreateCommunityVM struct {
	Name string `form:"name"`
}

type CommunityVM struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
