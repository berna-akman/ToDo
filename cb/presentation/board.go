package presentation

import (
	"github.com/google/uuid"
	"to-do-api/cb/domain/board"
	"to-do-api/cb/domain/user"
	"to-do-api/cb/infrastructure/mq"
	"to-do-api/internal/errors"
)

type BoardService interface {
	GetAllBoards() (*[]board.Board, error)
	GetBoardByID(string) (*board.Board, error)
	CreateBoard(board.Board) (*board.CreateResponse, error)
	UpdateBoard(string, board.Board) error
	DeleteBoard(string) error
	AddColumnToBoard(string, board.Column) (*board.CreateResponse, error)
	RemoveColumnFromBoard(board.DeleteColumnRequest) error
	CreateCard(string, board.CreateCardRequest) (*board.CreateResponse, error)
	GetCards(board.GetCardRequest) (*[]board.Card, error)
	CreateCardAssignee(board.CreateCardAssigneeRequest) error
}

var defaultColumns = []string{"To Do", "In Progress", "In Test", "Done"}

type boardService struct {
	boardRepository board.BoardRepository
	userRepository  user.UserRepository
	userMap         map[string]string
	// TODO: mq
}

func NewBoardService(boardRepository board.BoardRepository, userRepository user.UserRepository, userMap map[string]string) BoardService {
	return boardService{boardRepository: boardRepository, userRepository: userRepository, userMap: userMap}
}

func (s boardService) GetAllBoards() (*[]board.Board, error) {
	boards, err := s.boardRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return boards, nil
}

func (s boardService) GetBoardByID(id string) (*board.Board, error) {
	b, err := s.boardRepository.GetByID(id)
	if err != nil {
		return nil, errors.ErrorBoardNotFound
	}

	return b, nil
}

func (s boardService) CreateBoard(b board.Board) (*board.CreateResponse, error) {
	b.ID = uuid.NewString()
	if len(b.Columns) == 0 {
		// Set default columns
		b.Columns = make([]board.Column, len(defaultColumns))
		for i := range defaultColumns {
			column := board.Column{
				ColumnID: uuid.NewString(),
				Name:     defaultColumns[i],
				Cards:    make([]board.Card, 0),
			}
			b.Columns[i] = column
		}
	} else {
		for i := range b.Columns {
			b.Columns[i].ColumnID = uuid.NewString()
			b.Columns[i].Cards = make([]board.Card, 0)
		}
	}

	return s.boardRepository.CreateBoard(b)
}

func (s boardService) UpdateBoard(boardID string, b board.Board) error {
	return s.boardRepository.Update(boardID, b)
}

func (s boardService) DeleteBoard(id string) error {
	_, err := s.boardRepository.GetByID(id)
	if err != nil {
		return errors.ErrorBoardNotFound
	}
	return s.boardRepository.Delete(id)
}

func (s boardService) AddColumnToBoard(boardID string, column board.Column) (*board.CreateResponse, error) {
	column.ColumnID = uuid.NewString()
	column.Cards = make([]board.Card, 0)
	return s.boardRepository.AddColumnToBoard(boardID, column)
}

func (s boardService) RemoveColumnFromBoard(req board.DeleteColumnRequest) error {
	return s.boardRepository.RemoveColumnFromBoard(req)
}

func (s boardService) CreateCard(boardID string, req board.CreateCardRequest) (*board.CreateResponse, error) {
	req.Card.ID = uuid.NewString()
	return s.boardRepository.CreateCard(boardID, req)
}

func (s boardService) GetCards(req board.GetCardRequest) (*[]board.Card, error) {
	cards, err := s.boardRepository.GetCards(req)
	userMap := s.userMap
	var cardsWithUsers []board.Card
	// Get Cards if and only if assigneeId has a name map
	for _, v := range *cards {
		if len(userMap[v.AssigneeID]) > 0 {
			cardsWithUsers = append(cardsWithUsers, v)
			// TODO: bulamazsa db'ye gitsin
		}
	}
	return &cardsWithUsers, err
}

func (s boardService) CreateCardAssignee(req board.CreateCardAssigneeRequest) error {
	cardId, err := s.boardRepository.CreateCardAssignee(req)
	u, err := s.userRepository.GetByID(req.AssigneeID)
	err = s.userRepository.AddCardIdToUser(u.ID, cardId)

	if err != nil {
		return err
	}
	// TODO: produce message
	err = mq.ProduceMessage() // TODO: model, message return etsin err yerine
	if err != nil {
		return err
	}
	return err
}
