package main

import (
	"fmt"

	"github.com/Aarya-Patel/deck/pkg/deck"
)

func main() {
	deck, err := deck.New()
	if err != nil {
		panic(err)
	}

	fmt.Println(deck)
}
