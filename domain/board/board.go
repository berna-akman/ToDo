package board

type BoardRes struct {
	DTO
	Board
	Boards
}

type DTO struct {
	Board      Boards `json:"boards"`
	TotalCount int    `json:"totalCount"`
}

type Boards []Board

type Board struct {
	ID          int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      Status `json:"status" gorm:"foreignKey:ID"`
}

type Status struct {
	ID        uint32 `json:"id"`
	IsActive  bool   `json:"isActive" gorm:"default=true"` // neden default true olmadı?
	IsStarred bool   `json:"isStarred"`
}

type GetBoardRequest struct {
	ID int `json:"id"`
}

// TODO: board'a ait br logic'se burada don
