package presentation

import (
	"to-do-api/domain/board"
	"to-do-api/internal/errors"
)

type BoardService interface {
	GetAll() (*board.DTO, error)
	GetByID(uint) (*board.Board, error)
	Create(board.Board) error
	Update(board.Board) error
	Delete(uint) error
}

type boardService struct {
	r board.BoardRepository
}

func NewBoardService(repository board.BoardRepository) BoardService {
	return boardService{repository}
}

func (s boardService) GetAll() (*board.DTO, error) {
	boards, err := s.r.FindAll()
	if err != nil {
		return nil, err
	}

	return boards, nil
}

func (s boardService) GetByID(id uint) (*board.Board, error) {
	b, err := s.r.GetByID(id)
	if err != nil {
		return nil, errors.ErrorBoardNotFound
	}

	return b, nil
}

func (s boardService) Create(b board.Board) error {
	return s.r.Create(b)
}

func (s boardService) Update(b board.Board) error {
	_, err := s.r.GetByID(b.BoardID)
	if err != nil {
		return errors.ErrorBoardNotFound
	}

	return s.r.Update(b)
}

func (s boardService) Delete(id uint) error {
	_, err := s.r.GetByID(id)
	if err != nil {
		return errors.ErrorBoardNotFound
	}
	return s.r.Delete(id)
}
