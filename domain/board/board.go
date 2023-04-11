package board

import (
	"gorm.io/gorm"
	"time"
)

type DTO struct {
	Board []Board `json:"boards"`
	Count int64   `json:"count"`
}

type Board struct {
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	ID          uint           `json:"id" gorm:"primaryKey"`
	BoardID     uint           `json:"board_id" gorm:"index"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
}

type GetBoardRequest struct {
	ID int `json:"id"`
}

// TODO: board'a ait br logic'se burada don
