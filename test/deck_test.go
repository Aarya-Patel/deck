package test

import (
	"testing"

	"github.com/Aarya-Patel/deck/pkg/card"
	"github.com/Aarya-Patel/deck/pkg/deck"
)

func TestDefaultDeck(t *testing.T) {
	d := deck.New()
	if len(*d) != 52 {
		t.Error("Expected 52 cards, got", len(*d))
	}

	var suitCounts [4]int
	var rankCounts [13]int

	for _, c := range *d {
		if c.Suit == card.Joker {
			t.Error("Found a joker within the default deck")
		}
		if !card.IsValidSuit(c.Suit) {
			t.Error("Invalid suit", c.Suit.String())
		}
		if !card.IsValidRank(c.Rank) {
			t.Error("Invalid rank", c.Rank.String())
		}
		suitCounts[c.Suit-1] += 1
		rankCounts[c.Rank-1] += 1
	}

	for i, count := range suitCounts {
		if count != 13 {
			t.Error("Expected 13 cards of", card.Suits[i].String(), "\nGot", count)
		}
	}
	for i, count := range rankCounts {
		if count != 4 {
			t.Error("Expected 4 cards of", card.Ranks[i].String(), "\nGot", count)
		}
	}
}

func TestDeckWithJokers(t *testing.T) {
	const numJokers int = 3
	countJokers := 0

	d := deck.New(
		deck.WithJokers(numJokers),
	)

	for _, c := range *d {
		if c.Suit == card.Joker {
			countJokers++
		}
	}

	if countJokers != numJokers {
		t.Errorf("Expected %d jokers. Got %d", numJokers, countJokers)
	}
}

func TestCustomSort(t *testing.T) {
	d := deck.New(
		// Reverse sort on card.Rank
		deck.WithCustomSort(func(d *deck.Deck) func(int, int) bool {
			return func(i, j int) bool {
				return (*d)[i].Rank > (*d)[j].Rank
			}
		}),
	)

	prevRank := (*d)[0].Rank
	for i := 1; i < len(*d); i++ {
		c := (*d)[i]
		if prevRank < c.Rank {
			t.Errorf("Previous card rank of %s is smaller than current rank of %s", prevRank.String(), c.Rank.String())
		}
		prevRank = c.Rank
	}
}

func TestShuffle(t *testing.T) {
	d := deck.New(
		deck.WithShuffle(),
	)

	firstCard := (*d)[0]
	lastCard := (*d)[len(*d)-1]

	if firstCard.Suit == card.Spades && firstCard.Rank == card.Ace {
		t.Error("First card in the deck shouldn't be Ace of Spades")
	}
	if lastCard.Suit == card.Hearts && lastCard.Rank == card.King {
		t.Error("Last card in the deck shouldn't be King of Hearts")
	}
}

func TestWithFilter(t *testing.T) {
	d := deck.New(
		deck.WithFilter(func(c *card.Card) bool {
			return (c.Suit == card.Diamonds || c.Rank == card.Six)
		}),
	)
	for _, c := range *d {
		if c.Suit == card.Diamonds {
			t.Errorf("Expected no card with %s suit", card.Diamonds.String())
		}
		if c.Rank == card.Six {
			t.Errorf("Expected no card with rank %s", card.Six.String())
		}
	}
}

func TestWithAdditionalDecks(t *testing.T) {
	const numAdditionalDecks int = 1
	d := deck.New(
		deck.WithAdditionalDecks(numAdditionalDecks),
	)
	if len(*d) != 52+(numAdditionalDecks*52) {
		t.Errorf("Expected %d cards. Got %d", 52+(numAdditionalDecks*52), len(*d))
	}

	var suitCounts [4]int
	var rankCounts [13]int

	for _, c := range *d {
		if c.Suit == card.Joker {
			t.Error("Found a Joker within the default deck")
		}
		if !card.IsValidSuit(c.Suit) {
			t.Errorf("Invalid suit %s", c.Suit.String())
		}
		if !card.IsValidRank(c.Rank) {
			t.Errorf("Invalid rank %s", c.Rank.String())
		}
		suitCounts[c.Suit-1] += 1
		rankCounts[c.Rank-1] += 1
	}

	for i, count := range suitCounts {
		if count != 13+(numAdditionalDecks*13) {
			t.Error("Expected 13 cards of", card.Suits[i].String(), "\nGot", count)
		}
	}
	for i, count := range rankCounts {
		if count != 4+(numAdditionalDecks*4) {
			t.Error("Expected 4 cards of", card.Ranks[i].String(), "\nGot", count)
		}
	}
}
