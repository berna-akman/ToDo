package board

type BoardRepository interface {
	FindAll() (*[]Board, error)
	GetByID(string) (*Board, error)
	Create(Board) (*CreateResponse, error)
	Update(Board) error
	Delete(string) error
	AddColumn(b Board) ([]string, error)
}
