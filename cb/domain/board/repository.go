package board

type BoardRepository interface {
	FindAll() (*[]Board, error)
	GetByID(string) (*Board, error)
	CreateBoard(Board) (*CreateResponse, error)
	Update(string, Board) error
	Delete(string) error
	AddColumnToBoard(string, Column) (*CreateResponse, error)
	CreateCard(string, Card) (*CreateResponse, error)
	RemoveColumnFromBoard(string, string) error
	GetCards(string, string, string) (*[]Card, error)
}
