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

func (r BoardRepository) Update(boardID string, b board.Board) error {
	// TODO: How can I replace with only values from request body
	_, err := r.bucket.DefaultCollection().Replace(boardID, b, nil)
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

func (r BoardRepository) AddColumnToBoard(boardID string, column board.Column) (*board.CreateResponse, error) {
	// Append column to column array
	mops := []gocb.MutateInSpec{
		gocb.ArrayAppendSpec("columns", column, nil),
	}
	_, err := r.collection.MutateIn(boardID, mops, nil)
	if err != nil {
		return nil, err
	}

	return &board.CreateResponse{ID: column.ColumnID}, nil
}

func (r BoardRepository) RemoveColumnFromBoard(boardID string, columnID string) error {
	var doc board.Board
	result, err := r.collection.Get(boardID, nil)
	if err != nil {
		return err
	}

	if err = result.Content(&doc); err != nil {
		return err
	}

	var index int
	for i, v := range doc.Columns {
		if v.ColumnID == columnID {
			index = i
			break
		}
	}

	path := fmt.Sprintf("columns[%d]", index)
	mops := []gocb.MutateInSpec{
		gocb.RemoveSpec(path, nil),
	}
	_, err = r.collection.MutateIn(boardID, mops, nil)
	if err != nil {
		return err
	}

	return nil
}

func (r BoardRepository) CreateCard(boardID string, card board.Card) (*board.CreateResponse, error) {
	// Add to first column as default when creating new card
	mops := []gocb.MutateInSpec{
		gocb.ArrayAppendSpec("columns[0].cards", card, nil),
	}
	_, err := r.collection.MutateIn(boardID, mops, nil)
	if err != nil {
		return nil, err
	}

	return &board.CreateResponse{ID: card.ID}, nil
}

func (r BoardRepository) GetCards(boardID, columnID, assignee string) (*[]board.Card, error) {
	cards := make([]board.Card, 0)
	// Default query gets all cards from requested board
	query := fmt.Sprintf("SELECT card.id, card.summary, card.description, card.assignee FROM board b UNNEST b.columns column UNNEST column.cards card WHERE b.id = '%s' ", boardID)
	// Filter by columnID
	if len(columnID) > 0 {
		query = query + "AND column.columnId = $columnID "
	}
	// Filter by assignee
	if len(assignee) > 0 {
		query = query + "AND card.assignee = $assignee"
	}
	params := map[string]interface{}{
		"columnID": columnID,
		"assignee": assignee,
	}
	rows, err := r.cluster.Query(query, &gocb.QueryOptions{NamedParameters: params})
	if err != nil {
		return &[]board.Card{}, err
	}

	var c board.Card
	defer rows.Close()
	for rows.Next() {
		err = rows.Row(&c)
		if err != nil {
			return nil, err
		}
		cards = append(cards, c)
	}

	return &cards, nil
}
