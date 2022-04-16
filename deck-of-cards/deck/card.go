//go:generate stringer -type=CardRank,CardSuit
package deck

import "fmt"

type CardRank uint8
type CardSuit uint8

const (
	_ CardRank = iota
	Ace
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
	Diamond CardSuit = iota
	Club
	Heart
	Spade
	Joker
)

type Card struct {
	Rank CardRank
	Suit CardSuit
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}
