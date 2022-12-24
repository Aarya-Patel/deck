package main

import (
	"fmt"

	"github.com/Aarya-Patel/deck/pkg/card"
	"github.com/Aarya-Patel/deck/pkg/deck"
)

func main() {
	var d *deck.Deck

	d = deck.New(
		deck.WithCustomSort(func(d *deck.Deck) func(int, int) bool {
			return func(i, j int) bool {
				return (*d)[i].Rank > (*d)[j].Rank
			}
		}),
		deck.WithJokers(2),
		deck.WithShuffle(),
		deck.WithFilter(func(c *card.Card) bool {
			return c.Suit == card.Hearts
		}),
	)

	for _, card := range *d {
		fmt.Println(*card)
	}

}
