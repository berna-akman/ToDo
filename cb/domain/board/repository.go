package board

type BoardRepository interface {
	FindAll() (*[]Board, error)
	GetByID(string) (*Board, error)
	Create(Board) (string, error)
	Update(Board) error
	Delete(string) error
}
