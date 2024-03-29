package board

type BoardRepository interface {
	FindAll() (*[]Board, error)
	GetByID(string) (*Board, error)
	CreateBoard(Board) (*CreateResponse, error)
	Update(string, Board) error
	Delete(string) error
	AddColumnToBoard(string, Column) (*CreateResponse, error)
	CreateCard(string, CreateCardRequest) (*CreateResponse, error)
	RemoveColumnFromBoard(DeleteColumnRequest) error
	GetCards(GetCardRequest) (*[]Card, error)
	CreateCardAssignee(CreateCardAssigneeRequest) (string, error)
}
