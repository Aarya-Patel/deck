package deck

import (
	"math/rand"
	"sort"

	"github.com/Aarya-Patel/deck/pkg/card"
)

type Deck []*card.Card

func New(options ...func(*Deck)) *Deck {
	var deck Deck

	buildDefaultDeck(&deck)

	for _, option := range options {
		option(&deck)
	}

	return &deck
}

func buildDefaultDeck(deck *Deck) {
	for _, suit := range card.Suits {
		for _, rank := range card.Ranks {
			cfg := card.Config{Suit: suit, Rank: rank}
			card, err := card.New(cfg)
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

func WithShuffle() func(*Deck) {
	return func(deck *Deck) {
		rand.Shuffle(len(*deck), func(i, j int) {
			(*deck)[i], (*deck)[j] = (*deck)[j], (*deck)[i]
		})
	}
}

func WithJokers(n int) func(*Deck) {
	return func(deck *Deck) {
		cfg := card.Config{Suit: card.Joker, Rank: 0}
		for i := 0; i < n; i++ {
			card, err := card.New(cfg)
			if err != nil {
				panic(err)
			}
			*deck = append(*deck, card)
		}
	}
}

func WithFilter(fn func(*card.Card) bool) func(*Deck) {
	return func(deck *Deck) {
		var newDeck Deck

		for _, card := range *deck {
			shouldFilter := fn(card)
			if !shouldFilter {
				newDeck = append(newDeck, card)
			}
		}

		*deck = newDeck
	}
}
