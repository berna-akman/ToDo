package persistence

import (
	"gorm.io/gorm"
	card2 "to-do-api/pg/domain/card"
)

type CardRepository struct {
	db *gorm.DB
}

func NewCardRepository(ds *DataSource) card2.CardRepository {
	db := ds.DB.Table("cards").Session(&gorm.Session{NewDB: true})
	return &CardRepository{
		db: db,
	}
}

func (r *CardRepository) GetAll() (card2.Card, error) {
	return card2.Card{}, nil
}

func (r *CardRepository) Create(card card2.Card) error {
	return r.db.Create(&card).Error
}
