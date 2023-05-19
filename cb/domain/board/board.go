package board

type Board struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Column      []string `json:"columns"` // TODO: default tagi bulamadım?
}

type CreateResponse struct {
	ID string `json:"id"`
}
