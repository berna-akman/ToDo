package persistence

import (
	"fmt"
	"github.com/couchbase/gocb/v2"
	"github.com/pkg/errors"
	"to-do-api/cb/domain/user"
)

type UserRepository struct {
	cluster    *gocb.Cluster
	bucket     *gocb.Bucket
	collection *gocb.Collection
}

func NewUserRepository(ds *DataSource) user.UserRepository {
	cb := ds.Cluster.Bucket("users").DefaultCollection()
	return &UserRepository{
		cluster:    ds.Cluster,
		bucket:     cb.Bucket(),
		collection: cb.Bucket().DefaultCollection(),
	}
}

func (r UserRepository) FindAll() (*[]user.User, error) {
	var users []user.User

	rows, err := r.cluster.Query("SELECT users.* FROM users", nil)
	if err != nil {
		return &[]user.User{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var b user.User
		err = rows.Row(&b)
		if err != nil {
			return nil, err
		}
		users = append(users, b)
	}

	return &users, nil
}

func (r UserRepository) CreateUser(u user.User) (*user.CreateResponse, error) {
	// TODO: Auto increment
	_, err := r.collection.Upsert(u.ID, u, nil)
	if err != nil {
		fmt.Println("Failed to create document:", err)
		return nil, err
	}

	return &user.CreateResponse{ID: u.ID}, nil
}

func (r UserRepository) GetByID(id string) (*user.User, error) {
	var b user.User
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

func (r UserRepository) AddCardIdToUser(assigneeID string, cardID string) error {
	// Append cardId to cardIds array
	mops := []gocb.MutateInSpec{
		gocb.ArrayAppendSpec("cardIds", cardID, nil),
	}
	_, err := r.collection.MutateIn(assigneeID, mops, nil)
	if err != nil {
		return err
	}

	return nil
}
