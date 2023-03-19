package card

type CardRepository interface {
	GetAll() (Card, error)
	Create(Card) error
}
