package card

/* ----- Types ------*/
type Suit int
type Value int

/* ----- Constants ------*/
var Suits = [...]Suit{
	Spades,
	Diamonds,
	Clubs,
	Hearts,
}

var Values = [...]Value{
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
	Ace Value = iota + 1
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
)

/* ----- Struct ------*/
type Card struct {
	Suit    Suit
	Value   Value
	IsJoker bool
}

type Config struct {
	Suit    Suit
	Value   Value
	IsJoker bool
}
