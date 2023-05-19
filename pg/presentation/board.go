package presentation

import (
	"to-do-api/internal/errors"
	board2 "to-do-api/pg/domain/board"
)

type BoardService interface {
	GetAllBoards() (*board2.DTO, error)
	GetBoardByID(uint) (*board2.Board, error)
	CreateBoard(board2.Board) error
	UpdateBoard(board2.Board) error
	DeleteBoard(uint) error
}

type boardService struct {
	r board2.BoardRepository
}

func NewBoardService(repository board2.BoardRepository) BoardService {
	return boardService{repository}
}

func (s boardService) GetAllBoards() (*board2.DTO, error) {
	boards, err := s.r.FindAll()
	if err != nil {
		return nil, err
	}

	return boards, nil
}

func (s boardService) GetBoardByID(id uint) (*board2.Board, error) {
	b, err := s.r.GetByID(id)
	if err != nil {
		return nil, errors.ErrorBoardNotFound
	}

	return b, nil
}

func (s boardService) CreateBoard(b board2.Board) error {
	return s.r.Create(b)
}

func (s boardService) UpdateBoard(b board2.Board) error {
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
