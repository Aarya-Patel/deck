package card

import (
	"errors"
)

/* ----- Package Methods ------*/
func New(s Suit, v Value) (*Card, error) {
	if s < Spades || s > Hearts {
		return &Card{}, errors.New("suit: Invalid suit for card")
	}

	if v < Ace || v > King {
		return &Card{}, errors.New("value: Invalid value for card")
	}

	return &Card{s, v}, nil
}
