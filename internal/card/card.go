package card

import (
	"errors"
)

/* ----- Package Methods ------*/
func New(cfg Config) (*Card, error) {
	if cfg.IsJoker == false && (cfg.Suit == 0 || cfg.Value == 0) {
		return &Card{}, errors.New("config: Invalid config for New()")
	}

	if cfg.IsJoker {
		return &Card{IsJoker: true}, nil
	}

	if cfg.Suit < Spades || cfg.Suit > Hearts {
		return &Card{}, errors.New("suit: Invalid suit for card")
	}

	if cfg.Value < Ace || cfg.Value > King {
		return &Card{}, errors.New("value: Invalid value for card")
	}

	return &Card{Suit: cfg.Suit, Value: cfg.Value, IsJoker: cfg.IsJoker}, nil
}
