package board

type Board struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Columns     []Column `json:"columns"`
}

type Column struct {
	ColumnID string `json:"columnId"`
	Name     string `json:"name"`
	Cards    []Card `json:"cards"`
}

type Card struct {
	ID          string `json:"id"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
}

type CreateResponse struct {
	ID string `json:"id"`
}
