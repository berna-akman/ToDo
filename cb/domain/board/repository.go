package board

type BoardRepository interface {
	FindAll() (*[]Board, error)
	GetByID(string) (*Board, error)
	CreateBoard(Board) (*CreateResponse, error)
	Update(string, Board) error
	Delete(string) error
	CreateCard(Board) (*CreateResponse, error)
}
