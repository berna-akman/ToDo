package persistence

import (
	"gorm.io/gorm"
	"to-do-api/domain/card"
)

type CardRepository struct {
	db *gorm.DB
}

func NewCardRepository(ds *DataSource) card.CardRepository {
	db := ds.DB.Table("cards").Session(&gorm.Session{NewDB: true})
	return &CardRepository{
		db: db,
	}
}

func (r *CardRepository) GetAll() (card.Card, error) {
	return card.Card{}, nil
}

func (r *CardRepository) Create(card card.Card) error {
	return r.db.Create(&card).Error
}
