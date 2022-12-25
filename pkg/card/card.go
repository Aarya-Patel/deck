package card

import (
	"errors"
)

/* ----- Package Methods ------*/
func New(cfg Config) (*Card, error) {
	if cfg.Suit == Joker {
		return &Card{Suit: cfg.Suit, Rank: 0}, nil
	}

	if !IsValidSuit(cfg.Suit) || !IsValidRank(cfg.Rank) {
		return &Card{}, errors.New("config: Invalid config")
	}

	if cfg.Rank < Ace || cfg.Rank > King {
		return &Card{}, errors.New("rank: Invalid Rank for card")
	}

	return &Card{Suit: cfg.Suit, Rank: cfg.Rank}, nil
}

func IsValidSuit(s Suit) bool {
	return (Spades <= s && s <= Joker)
}

func IsValidRank(r Rank) bool {
	return (Ace <= r && r <= King)
}
