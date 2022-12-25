package test

import (
	"testing"

	"github.com/Aarya-Patel/deck/pkg/card"
)

func TestValidCards(t *testing.T) {
	for _, suit := range card.Suits {
		for _, rank := range card.Ranks {
			cfg := card.Config{Suit: suit, Rank: rank}
			c, err := card.New(cfg)

			if c.Suit != suit ||
				c.Rank != rank ||
				err != nil {
				t.Error("creating default cards failed")
			}
		}
	}
}

func TestCreateJokers(t *testing.T) {
	cfgs := []card.Config{
		{Suit: card.Joker, Rank: 0},
		{Suit: card.Joker, Rank: 1},
	}

	for _, cfg := range cfgs {
		c, err := card.New(cfg)
		if err != nil {
			t.Error("creating jokers failed due to error:", err)
		}

		if c.Suit != card.Joker || c.Rank != 0 {
			t.Error("Expected {Suit: Joker, Rank: 0}, Got:", c)
		}
	}
}

func TestInvalidConfig(t *testing.T) {
	cfgs := []card.Config{
		{Suit: card.Hearts, Rank: 0},
		{Suit: 0, Rank: 1},
		{Suit: 0, Rank: 0},
	}

	for _, cfg := range cfgs {
		c, err := card.New(cfg)
		if err == nil {
			t.Error("invalid card created", c)
		}
	}

}
