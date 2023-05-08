package presentation

import (
	"to-do-api/domain/board"
	"to-do-api/internal/errors"
)

type BoardService interface {
	GetAllBoards() (*board.DTO, error)
	GetBoardByID(uint) (*board.Board, error)
	CreateBoard(board.Board) error
	UpdateBoard(board.Board) error
	DeleteBoard(uint) error
}

type boardService struct {
	r board.BoardRepository
}

func NewBoardService(repository board.BoardRepository) BoardService {
	return boardService{repository}
}

func (s boardService) GetAllBoards() (*board.DTO, error) {
	boards, err := s.r.FindAll()
	if err != nil {
		return nil, err
	}

	return boards, nil
}

func (s boardService) GetBoardByID(id uint) (*board.Board, error) {
	b, err := s.r.GetByID(id)
	if err != nil {
		return nil, errors.ErrorBoardNotFound
	}

	return b, nil
}

func (s boardService) CreateBoard(b board.Board) error {
	return s.r.Create(b)
}

func (s boardService) UpdateBoard(b board.Board) error {
	_, err := s.r.GetByID(b.BoardID)
	if err != nil {
		return errors.ErrorBoardNotFound
	}

	return s.r.Update(b)
}

func (s boardService) DeleteBoard(id uint) error {
	_, err := s.r.GetByID(id)
	if err != nil {
		return errors.ErrorBoardNotFound
	}
	return s.r.Delete(id)
}
