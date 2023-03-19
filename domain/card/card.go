package card

import (
	"github.com/google/uuid"
	"time"
)

type Card struct {
	ID               uuid.UUID  `json:"id"`
	Name             string     `json:"name"`
	Description      string     `json:"description"`
	Status           string     `json:"status"`
	Closed           bool       `json:"closed"` // card can have closed: false but be on an archived board
	LastActivityDate *time.Time `json:"lastActivityDate"`
	Due              *time.Time `json:"due"`     // due date on the card, if exists
	DueDate          bool       `json:"dueDate"` // whether the due date has been marked complete
	BoardID          int        `json:"boardId"`
	//Labels
	//Lists
}

type GetCardRequest struct {
	ID uint `json:"id"`
}
