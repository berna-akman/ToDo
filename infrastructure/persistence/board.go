package persistence

import (
	"gorm.io/gorm"
	"to-do-api/domain/board"
)

type BoardRepository struct {
	db *gorm.DB
}

func NewBoardRepository(ds *DataSource) board.BoardRepository {
	db := ds.DB.Table("boards").Session(&gorm.Session{NewDB: true})
	return &BoardRepository{
		db: db,
	}
}

func (r *BoardRepository) FindAll() (board.Boards, error) {
	// TODO: repo'dan count alabilirsin
	var b board.Boards
	return b, r.db.Find(&b).Error
}

func (r *BoardRepository) GetByID(id int) (*board.Board, error) {
	var b board.Board
	return &b, r.db.Where("id = ?", id).First(&b).Error
}

func (r *BoardRepository) Create(b board.Board) error {
	return r.db.Create(&b).Error
}

func (r *BoardRepository) Update(b board.Board) error {
	return r.db.Model(&b).Updates(board.Board{ID: b.ID, Name: b.Name, Description: b.Description, Status: b.Status}).Error
}

func (r *BoardRepository) Delete(id int) error {
	return r.db.Delete(board.Board{}, "id = ?", id).Error
}
