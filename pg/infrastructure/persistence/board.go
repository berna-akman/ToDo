package persistence

import (
	"gorm.io/gorm"
	board2 "to-do-api/pg/domain/board"
)

type BoardRepository struct {
	db *gorm.DB
}

func NewBoardRepository(ds *DataSource) board2.BoardRepository {
	db := ds.DB.Table("boards").Session(&gorm.Session{NewDB: true})
	return &BoardRepository{
		db: db,
	}
}

func (r BoardRepository) FindAll() (*board2.DTO, error) {
	var boards []board2.Board
	var count int64

	err := r.db.Find(&boards).Error
	if err != nil {
		return &board2.DTO{}, err
	}

	err = r.db.Model(&board2.Board{}).Count(&count).Error
	if err != nil {
		return &board2.DTO{}, err
	}
	return &board2.DTO{
		Board: boards,
		Count: count,
	}, nil
}

func (r BoardRepository) GetByID(id uint) (*board2.Board, error) {
	var b board2.Board
	return &b, r.db.Where("board_id = ?", id).First(&b).Error
}

func (r BoardRepository) Create(b board2.Board) error {
	return r.db.Create(&b).Error
}

func (r BoardRepository) Update(b board2.Board) error {
	return r.db.Model(&b).Where("board_id = ?", b.BoardID).Updates(board2.Board{ID: b.ID, BoardID: b.BoardID, Name: b.Name, Description: b.Description}).Error
}

func (r BoardRepository) Delete(id uint) error {
	return r.db.Where("board_id = ?", id).Delete(&board2.Board{}).Error
}
