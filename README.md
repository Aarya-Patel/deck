# Deck of Cards
This is a package that represents a deck of cards. Available are 2 public packages:
- `deck` - Represents the deck of cards; provides functionality on the deck level.
- `card` - Represents the cards within the deck.

## Package `deck`
Here are the list of package methods for `deck`:
- `New() *Deck` - creates a new deck of cards. `New` can take options to help customize the deck. Here are all the functional options that can be passed into `New`:
  - `WithCustomSort(fn func(*Deck) func(int, int) bool)` - provides custom sorting. Based off of the `sort` package, `fn` uses the `*Deck` closure to return a comparator function of `func(int, int) bool`.
  - `WithShuffle()` - shuffles the deck.
  - `WithJokers(n int)` - adds `n` cards of type `card.Joker` to the deck.
  - `WithFilter(fn func(*card.Card) bool)` - filters out cards in the deck as specified by the return value of `fn`. 
  - `WithAdditionalDecks(n int)` - create `n` additional copies of the **current** modified deck.

## Package `card`
Here are the list of package methods for `card`:
- `New(cfg Config) (*Card, error)` - creates a `card.Card` with the specified outlined by `cfg`.
- `IsValidSuit(s Suit) bool` - checks if `s` is a valid `Suit`
- `IsValidRank(s Rank) bool` - checks if `r` is a valid `Rank`
