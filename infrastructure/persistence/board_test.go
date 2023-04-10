package persistence

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"to-do-api/domain/board"
)

var ds *DataSource

type BoardRepositoryTestSuite struct {
	suite.Suite
	r board.BoardRepository
}

func Test_RunBoardRepositoryTestSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("board repository integration tests skipped")
	}

	suite.Run(t, new(BoardRepositoryTestSuite))
}

func (r *BoardRepositoryTestSuite) SetupSuite() {
	r.NoError(ds.DB.AutoMigrate(&board.Board{}))
	r.r = NewBoardRepository(ds)
}

func (r *BoardRepositoryTestSuite) TestBoardRepository_GetByID() {
	data := board.Board{
		ID:          11,
		Name:        "test board name 11",
		Description: "test board desc 11",
		Status: board.Status{
			ID:        1,
			IsActive:  true,
			IsStarred: true,
		},
	}
	r.NoError(r.r.Create(data))

	actual, _ := r.r.GetByID(11)
	r.Equal(data, actual)
}

//func TestBoardRepository_Create(t *testing.T) {
//	type fields struct {
//		db *gorm.DB
//	}
//	type args struct {
//		b board.Board
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			r := BoardRepository{
//				db: tt.fields.db,
//			}
//			if err := r.Create(tt.args.b); (err != nil) != tt.wantErr {
//				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func TestBoardRepository_Delete(t *testing.T) {
//	type fields struct {
//		db *gorm.DB
//	}
//	type args struct {
//		id int
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			r := BoardRepository{
//				db: tt.fields.db,
//			}
//			if err := r.Delete(tt.args.id); (err != nil) != tt.wantErr {
//				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func TestBoardRepository_FindAll(t *testing.T) {
//	type fields struct {
//		db *gorm.DB
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		want    board.Boards
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			r := BoardRepository{
//				db: tt.fields.db,
//			}
//			got, err := r.FindAll()
//			if (err != nil) != tt.wantErr {
//				t.Errorf("FindAll() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("FindAll() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestBoardRepository_GetByID(t *testing.T) {
//	type fields struct {
//		db *gorm.DB
//	}
//	type args struct {
//		id int
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *board.Board
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			r := BoardRepository{
//				db: tt.fields.db,
//			}
//			got, err := r.GetByID(tt.args.id)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetByID() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestBoardRepository_Update(t *testing.T) {
//	type fields struct {
//		db *gorm.DB
//	}
//	type args struct {
//		b board.Board
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			r := BoardRepository{
//				db: tt.fields.db,
//			}
//			if err := r.Update(tt.args.b); (err != nil) != tt.wantErr {
//				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func TestNewBoardRepository(t *testing.T) {
//	type args struct {
//		ds *DataSource
//	}
//	tests := []struct {
//		name string
//		args args
//		want board.BoardRepository
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewBoardRepository(tt.args.ds); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewBoardRepository() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
