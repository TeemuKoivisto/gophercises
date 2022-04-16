//go:generate stringer -type=CardRank,CardSuit
package deck

import (
	"fmt"
	"sort"
)

type CardRank uint8

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
	minRank = Ace
	maxRank = King
)

type CardSuit uint8

const (
	Diamond CardSuit = iota
	Club
	Heart
	Spade
	Joker
)

var suits = [...]CardSuit{Diamond, Club, Heart, Spade}

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

func New() []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}
	return cards
}

type BySuitAndRank []Card

func (a BySuitAndRank) Len() int           { return len(a) }
func (a BySuitAndRank) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BySuitAndRank) Less(i, j int) bool { return absRank(a[i]) < absRank(a[j]) }

func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

func RegularSort(cards []Card) []Card {
	sort.Sort(BySuitAndRank(cards))
	return cards
}
