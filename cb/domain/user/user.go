package user

type User struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	CardIDs []string `json:"cardIds"`
}

type CreateResponse struct {
	ID string `json:"id"`
}
