package presentation

import (
	"github.com/google/uuid"
	"to-do-api/cb/domain/board"
	"to-do-api/internal/errors"
)

type BoardService interface {
	GetAllBoards() (*[]board.Board, error)
	GetBoardByID(string) (*board.Board, error)
	CreateBoard(board.Board) (*board.CreateResponse, error)
	UpdateBoard(string, board.Board) error
	DeleteBoard(string) error
	AddColumnToBoard(string, board.Column) (*board.CreateResponse, error)
	RemoveColumnFromBoard(string, string) error
	CreateCard(string, board.Card) (*board.CreateResponse, error)
	GetCardsByColumn(string, string) (*[]board.Card, error)
}

var defaultColumns = []string{"To Do", "In Progress", "In Test", "Done"}

type boardService struct {
	r board.BoardRepository
}

func NewBoardService(repository board.BoardRepository) BoardService {
	return boardService{repository}
}

func (s boardService) GetAllBoards() (*[]board.Board, error) {
	boards, err := s.r.FindAll()
	if err != nil {
		return nil, err
	}

	return boards, nil
}

func (s boardService) GetBoardByID(id string) (*board.Board, error) {
	b, err := s.r.GetByID(id)
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

	return s.r.CreateBoard(b)
}

func (s boardService) UpdateBoard(boardID string, b board.Board) error {
	return s.r.Update(boardID, b)
}

func (s boardService) DeleteBoard(id string) error {
	_, err := s.r.GetByID(id)
	if err != nil {
		return errors.ErrorBoardNotFound
	}
	return s.r.Delete(id)
}

func (s boardService) AddColumnToBoard(boardID string, column board.Column) (*board.CreateResponse, error) {
	column.ColumnID = uuid.NewString()
	column.Cards = make([]board.Card, 0)
	return s.r.AddColumnToBoard(boardID, column)
}

func (s boardService) RemoveColumnFromBoard(boardID string, columnID string) error {
	return s.r.RemoveColumnFromBoard(boardID, columnID)
}

func (s boardService) CreateCard(boardID string, card board.Card) (*board.CreateResponse, error) {
	card.ID = uuid.NewString()
	return s.r.CreateCard(boardID, card)
}

func (s boardService) GetCardsByColumn(boardID, columnID string) (*[]board.Card, error) {
	return s.r.GetCardsByColumn(boardID, columnID)
}
