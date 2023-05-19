package board

type Board struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Column      []string `json:"columns"`
	Status      string   `json:"status"`
}

type CreateResponse struct {
	ID string `json:"id"`
}
