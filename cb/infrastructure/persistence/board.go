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

func (r BoardRepository) RemoveColumnFromBoard(req board.DeleteColumnRequest) error {
	var doc board.Board
	result, err := r.collection.Get(req.BoardID, nil)
	if err != nil {
		return err
	}

	if err = result.Content(&doc); err != nil {
		return err
	}

	var index int
	for i, v := range doc.Columns {
		if v.ColumnID == req.ColumnID {
			index = i
			break
		}
	}

	path := fmt.Sprintf("columns[%d]", index)
	mops := []gocb.MutateInSpec{
		gocb.RemoveSpec(path, nil),
	}
	_, err = r.collection.MutateIn(req.BoardID, mops, nil)
	if err != nil {
		return err
	}

	return nil
}

func (r BoardRepository) CreateCard(boardID string, req board.CreateCardRequest) (*board.CreateResponse, error) {
	var doc board.Board
	result, err := r.collection.Get(boardID, nil)
	if err != nil {
		return nil, err
	}

	if err = result.Content(&doc); err != nil {
		return nil, err
	}

	var index int
	for i, v := range doc.Columns {
		if v.ColumnID == req.ColumnID {
			index = i
			break
		}
	}

	// Add to first column as default when creating new card
	path := "columns[0].cards"
	if len(req.ColumnID) > 0 {
		path = fmt.Sprintf("columns[%d].cards", index)
	}

	mops := []gocb.MutateInSpec{
		gocb.ArrayAppendSpec(path, req.Card, nil),
	}

	_, err = r.collection.MutateIn(boardID, mops, nil)
	if err != nil {
		return nil, err
	}

	return &board.CreateResponse{ID: req.Card.ID}, nil
}

func (r BoardRepository) GetCards(req board.GetCardRequest) (*[]board.Card, error) {
	cards := make([]board.Card, 0)
	// Default query gets all cards from requested board
	query := fmt.Sprintf("SELECT card.id, card.summary, card.description, card.assignee FROM board b UNNEST b.columns column UNNEST column.cards card WHERE b.id = '%s' ", req.BoardID)

	// Filter by columnID
	if len(req.ColumnID) > 0 {
		query = query + "AND column.columnId = $columnID "
	}

	// Filter by assignee
	if len(req.Assignee) > 0 {
		query = query + "AND card.assignee = $assignee"
	}

	params := map[string]interface{}{
		"columnID": req.ColumnID,
		"assignee": req.Assignee,
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

func (r BoardRepository) CreateCardAssignee(boardID, cardID string, req board.CreateCardAssigneeRequest) error {
	var doc board.Board
	result, err := r.collection.Get(boardID, nil)
	if err != nil {
		return err
	}

	if err = result.Content(&doc); err != nil {
		return err
	}

	var cardIndex, colIndex int
	var found bool
	for i, v1 := range doc.Columns {
		if found {
			break
		}
		for j, v2 := range v1.Cards {
			if v2.ID == cardID {
				colIndex = i
				cardIndex = j
				found = true
				break
			}
		}
	}

	path := fmt.Sprintf("columns[%d].cards[%d].assignee", colIndex, cardIndex)

	mops := []gocb.MutateInSpec{
		gocb.UpsertSpec(path, req.Assignee, nil),
	}

	_, err = r.collection.MutateIn(boardID, mops, nil)
	if err != nil {
		return err
	}

	return nil
}
