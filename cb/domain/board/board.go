package board

type Board struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Column      []string `json:"columns"`
	Card        []Card   `json:"cards"`
}

type Card struct {
	ID          string `json:"id"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type CreateResponse struct {
	ID string `json:"id"`
}
