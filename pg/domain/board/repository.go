package board

type BoardRepository interface {
	FindAll() (*DTO, error)
	GetByID(uint) (*Board, error)
	Create(Board) error
	Update(Board) error
	Delete(uint) error
}
