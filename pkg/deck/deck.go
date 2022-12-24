package deck

import (
	"math/rand"
	"sort"

	"github.com/Aarya-Patel/deck/pkg/card"
)

type Deck []*card.Card
type Option func(*Deck)

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

func WithCustomSort(fn func(*Deck) func(int, int) bool) Option {
	return func(deck *Deck) {
		closuredCustomSortFn := fn(deck)
		sort.Slice(*deck, closuredCustomSortFn)
	}
}

func WithShuffle() Option {
	return func(deck *Deck) {
		rand.Shuffle(len(*deck), func(i, j int) {
			(*deck)[i], (*deck)[j] = (*deck)[j], (*deck)[i]
		})
	}
}

func WithJokers(n int) Option {
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

func WithFilter(fn func(*card.Card) bool) Option {
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

func WithAdditionalDecks(n int) Option {
	return func(deck *Deck) {
		currentDeck := make(Deck, len(*deck))
		copy(currentDeck, *deck)

		for i := 0; i < n; i++ {
			for _, card := range currentDeck {
				*deck = append(*deck, card)
			}
		}
	}
}
