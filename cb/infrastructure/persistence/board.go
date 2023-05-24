package persistence

import (
	"fmt"
	"github.com/couchbase/gocb/v2"
	"github.com/pkg/errors"
	"to-do-api/cb/domain/board"
)

type BoardRepository struct {
	cluster    *gocb.Cluster
	bucket     *gocb.Bucket
	collection *gocb.Collection
}

func NewBoardRepository(ds *DataSource) board.BoardRepository {
	cb := ds.Cluster.Bucket("board").DefaultCollection()
	return &BoardRepository{
		cluster:    ds.Cluster,
		bucket:     cb.Bucket(),
		collection: cb.Bucket().DefaultCollection(),
	}
}

func (r BoardRepository) FindAll() (*[]board.Board, error) {
	var boards []board.Board

	rows, err := r.cluster.Query("SELECT board.* FROM board", nil)
	if err != nil {
		return &[]board.Board{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var b board.Board
		err = rows.Row(&b)
		if err != nil {
			return nil, err
		}
		boards = append(boards, b)
	}

	return &boards, nil
}

func (r BoardRepository) GetByID(id string) (*board.Board, error) {
	var b board.Board
	result, err := r.bucket.DefaultCollection().Get(id, nil)
	if err != nil {
		if errors.Is(err, gocb.ErrDocumentNotFound) {
			return nil, fmt.Errorf("document with ID '%s' not found", id)
		}
		return nil, err
	}

	err = result.Content(&b)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func (r BoardRepository) CreateBoard(b board.Board) (*board.CreateResponse, error) {
	_, err := r.collection.Upsert(b.ID, b, nil)
	if err != nil {
		fmt.Println("Failed to create document:", err)
		return nil, err
	}

	return &board.CreateResponse{ID: b.ID}, nil
}

func (r BoardRepository) Update(b board.Board) error {
	board := board.Board{
		ID:          b.ID,
		Name:        b.Name,
		Description: b.Description,
	}

	_, err := r.bucket.DefaultCollection().Upsert(b.ID, board, nil)
	if err != nil {
		return err
	}

	return nil
}

func (r BoardRepository) Delete(id string) error {
	_, err := r.bucket.DefaultCollection().Remove(id, nil)
	if err != nil {
		if errors.Is(err, gocb.ErrDocumentNotFound) {
			return fmt.Errorf("document with ID '%s' not found", id)
		}
		return err
	}

	return nil
}

func (r BoardRepository) CreateCard(b board.Board) (*board.CreateResponse, error) {
	_, err := r.collection.Upsert(b.ID, b, nil)
	if err != nil {
		fmt.Println("Failed to create document:", err)
		return nil, err
	}

	return &board.CreateResponse{ID: b.Card[len(b.Card)-1].ID}, nil
}