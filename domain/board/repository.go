package board

type BoardRepository interface {
	FindAll() (*DTO, error)
	GetByID(int) (*Board, error)
	Create(Board) error
	Update(Board) error
	Delete(int) error
}
