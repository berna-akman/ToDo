package presentation

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
	"to-do-api/domain/board"
	"to-do-api/infrastructure/persistence"
)

/*func Test_boardService_GetBoards(t *testing.T) {
	pg, mock, err := persistence.InitForTest()
	assert.Nil(t, err)
	repository := persistence.NewBoardRepository(pg)
	tt := time.Unix(1679923648, 0).UTC()

	columns := []string{"created_at", "updated_at", "deleted_at", "id", "board_id", "name", "description"}
	mock.ExpectQuery(`^SELECT(.+)FROM(.+)(.+)boards`).WillReturnRows(sqlmock.NewRows(columns).
		AddRow(tt, tt, nil, 1, 3, "fancy board name", "fancy board description"))

	s := &boardService{
		repository,
	}

	want := &board.DTO{
		Board: []board.Board{
			{
				CreatedAt:   tt,
				UpdatedAt:   tt,
				DeletedAt:   gorm.DeletedAt{},
				ID:          1,
				BoardID:     3,
				Name:        "fancy board name",
				Description: "fancy board description",
			},
		},

		Count: 1,
	}

	var got *board.DTO
	// TODO: Fix count problem
	got, err = s.GetAll()
	if err != nil {
		panic(err)
	}
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}*/

func Test_boardService_GetByID(t *testing.T) {
	pg, mock, err := persistence.InitForTest()
	assert.Nil(t, err)
	repository := persistence.NewBoardRepository(pg)
	tt := time.Unix(1679923648, 0).UTC()

	columns := []string{"created_at", "updated_at", "deleted_at", "id", "board_id", "name", "description"}
	mock.ExpectQuery(`^SELECT(.+)FROM(.+)(.+)boards`).WillReturnRows(sqlmock.NewRows(columns).
		AddRow(tt, tt, nil, 2, 4, "get board name", "get board description"))

	s := &boardService{
		repository,
	}

	want := &board.Board{
		CreatedAt:   tt,
		UpdatedAt:   tt,
		DeletedAt:   gorm.DeletedAt{},
		ID:          2,
		BoardID:     4,
		Name:        "get board name",
		Description: "get board description",
	}

	var got *board.Board
	got, err = s.GetBoardByID(4)
	if err != nil {
		panic(err)
	}
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}
