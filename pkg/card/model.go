package card

/* ----- Types ------*/
type Suit int
type Rank int

/* ----- Constants ------*/
var Suits = [...]Suit{
	Spades,
	Diamonds,
	Clubs,
	Hearts,
}

var Ranks = [...]Rank{
	Ace,
	Two,
	Three,
	Four,
	Five,
	Six,
	Seven,
	Eight,
	Nine,
	Ten,
	Jack,
	Queen,
	King,
}

/* ----- Enums ------*/
const (
	Ace Rank = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	Spades Suit = iota + 1
	Diamonds
	Clubs
	Hearts
	Joker
)

/* ----- Struct ------*/
type Card struct {
	Suit Suit
	Rank Rank
}

type Config struct {
	Suit Suit
	Rank Rank
}
