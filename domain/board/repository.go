package board

type BoardRepository interface {
	FindAll() (Boards, error)
	GetByID(int) (*Board, error)
	Create(Board) error
	Update(Board) error
	Delete(int) error
}
