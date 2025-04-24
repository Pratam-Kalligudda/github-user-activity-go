package model

type Event struct {
	ID      string  `json:"id"`
	Type    string  `json:"type"`
	Repo    Repo    `json:"repo"`
	Payload Payload `json:"payload"`
	Actor   Actor   `json:"actor"`
}

type Actor struct {
	Display_Login string `json:"display_login"`
}

type Repo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Payload struct {
	Repo_id      *int          `json:"repository_id"`
	Ref          *string       `json:"ref"`
	Ref_Type     *string       `json:"ref_type"`
	Pull_Request *Pull_Request `json:"pull_request"`
	Action       *string       `json:"action"`
	Commits      *[]Commit     `json:"commits"`
}
type Commit struct {
	Message string `json:"message"`
}

type Pull_Request struct {
	Title string `json:"title"`
	User  User   `user:"user"`
}

type User struct {
	Login string `json:"login"`
}
