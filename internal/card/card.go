package card

import (
	"errors"
)

/* ----- Package Methods ------*/
func New(cfg Config) (*Card, error) {
	if cfg.Suit != Joker && cfg.Rank == 0 {
		return &Card{}, errors.New("config: Invalid config for New()")
	}

	if cfg.Suit == Joker {
		return &Card{Suit: cfg.Suit, Rank: cfg.Rank}, nil
	}

	if cfg.Suit < Spades || cfg.Suit > Joker {
		return &Card{}, errors.New("suit: Invalid suit for card")
	}

	if cfg.Rank < Ace || cfg.Rank > King {
		return &Card{}, errors.New("rank: Invalid Rank for card")
	}

	return &Card{Suit: cfg.Suit, Rank: cfg.Rank}, nil
}
