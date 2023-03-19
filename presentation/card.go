package presentation

import (
	"to-do-api/domain/card"
)

type CardService interface {
	GetAll() (card.Card, error)
	Create(card.Card) error
}

type cardService struct {
	cardRepository card.CardRepository
}

func NewCardService(
	cardRepository card.CardRepository) CardService {
	return &cardService{
		cardRepository: cardRepository,
	}
}

func (s *cardService) GetAll() (card.Card, error) {
	return s.cardRepository.GetAll()
}

func (s *cardService) Create(card card.Card) error {
	return s.cardRepository.Create(card)
}
