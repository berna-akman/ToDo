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
	AssigneeID  string `json:"assigneeId"`
	Email       string `json:"email"`
}

type CreateResponse struct {
	ID string `json:"id"`
}

type CreateCardRequest struct {
	Card     Card   `json:"card"`
	ColumnID string `json:"columnId"`
}

type CreateCardAssigneeRequest struct {
	BoardID    string `json:"id"`
	CardID     string `json:"cardId"`
	AssigneeID string `json:"assigneeId"`
}

type GetCardRequest struct {
	BoardID    string `json:"id"`
	ColumnID   string `json:"columnId"`
	AssigneeID string `json:"assigneeId"`
}

type DeleteColumnRequest struct {
	BoardID  string `json:"id"`
	ColumnID string `json:"columnId"`
}
