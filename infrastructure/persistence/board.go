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

func (r BoardRepository) FindAll() (*board.DTO, error) {
	var boards []board.Board
	var count int64

	err := r.db.Find(&boards).Error
	if err != nil {
		return &board.DTO{}, err
	}

	err = r.db.Model(&board.Board{}).Count(&count).Error
	if err != nil {
		return &board.DTO{}, err
	}
	return &board.DTO{
		Board: boards,
		Count: count,
	}, nil
}

func (r BoardRepository) GetByID(id uint) (*board.Board, error) {
	var b board.Board
	return &b, r.db.Where("board_id = ?", id).First(&b).Error
}

func (r BoardRepository) Create(b board.Board) error {
	return r.db.Create(&b).Error
}

func (r BoardRepository) Update(b board.Board) error {
	return r.db.Model(&b).Where("board_id = ?", b.BoardID).Updates(board.Board{ID: b.ID, BoardID: b.BoardID, Name: b.Name, Description: b.Description}).Error
}

func (r BoardRepository) Delete(id uint) error {
	return r.db.Where("board_id = ?", id).Delete(&board.Board{}).Error
}
