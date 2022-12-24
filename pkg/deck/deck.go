package deck

import (
	"github.com/Aarya-Patel/deck/internal/card"
)

func New() ([]card.Card, error) {
	deck := []card.Card{}

	for _, suit := range card.Suits {
		for _, value := range card.Values {
			card, err := card.New(suit, value)
			if err != nil {
				return deck, err
			}
			deck = append(deck, card)
		}
	}

	return deck, nil
}
