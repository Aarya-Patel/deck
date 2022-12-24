package deck

import (
	"sort"

	"github.com/Aarya-Patel/deck/internal/card"
)

type Deck []*card.Card

func New(options ...func(*Deck)) *Deck {
	var deck Deck

	defaultOption(&deck)

	for _, option := range options {
		option(&deck)
	}

	return &deck
}

func defaultOption(deck *Deck) {
	for _, suit := range card.Suits {
		for _, value := range card.Values {
			card, err := card.New(suit, value)
			if err != nil {
				panic(err)
			}
			*deck = append(*deck, card)
		}
	}
}

func WithCustomSort(fn func(*Deck) func(int, int) bool) func(*Deck) {
	return func(deck *Deck) {
		closuredCustomSortFn := fn(deck)
		sort.Slice(*deck, closuredCustomSortFn)
	}
}
