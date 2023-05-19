package presentation

import (
	card2 "to-do-api/pg/domain/card"
)

type CardService interface {
	GetAll() (card2.Card, error)
	Create(card2.Card) error
}

type cardService struct {
	cardRepository card2.CardRepository
}

func NewCardService(
	cardRepository card2.CardRepository) CardService {
	return &cardService{
		cardRepository: cardRepository,
	}
}

func (s *cardService) GetAll() (card2.Card, error) {
	return s.cardRepository.GetAll()
}

func (s *cardService) Create(card card2.Card) error {
	return s.cardRepository.Create(card)
}
