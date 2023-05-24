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
	UpdateBoard(board.Board) error
	DeleteBoard(string) error
	CreateCard(board.Board) (*board.CreateResponse, error)
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
	if len(b.Column) == 0 {
		b.Column = defaultColumns
	}

	return s.r.CreateBoard(b)
}

func (s boardService) UpdateBoard(b board.Board) error {
	// TODO: Accept id from path when updating
	_, err := s.r.GetByID(b.ID)
	if err != nil {
		return errors.ErrorBoardNotFound
	}

	return s.r.Update(b)
}

func (s boardService) DeleteBoard(id string) error {
	_, err := s.r.GetByID(id)
	if err != nil {
		return errors.ErrorBoardNotFound
	}
	return s.r.Delete(id)
}

func (s boardService) CreateCard(b board.Board) (*board.CreateResponse, error) {
	return s.r.CreateCard(b)
}
