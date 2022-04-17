//go:generate stringer -type=CardRank,CardSuit
package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
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

func New(opts ...func([]Card) []Card) []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}
	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards
}

func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

type BySuitAndRank []Card

func (a BySuitAndRank) Len() int           { return len(a) }
func (a BySuitAndRank) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BySuitAndRank) Less(i, j int) bool { return absRank(a[i]) < absRank(a[j]) }

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

func RegularSort(cards []Card) []Card {
	sort.Sort(BySuitAndRank(cards))
	return cards
}

func Shuffle(cards []Card) []Card {
	ret := make([]Card, len(cards))
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i, j := range r.Perm(len(cards)) {
		ret[i] = cards[j]
	}
	return ret
}

func Jokers(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{
				Rank: CardRank(i),
				Suit: Joker,
			})
		}
		return cards
	}
}

func Filter(f func(card Card) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card
		for _, c := range cards {
			if !f(c) {
				ret = append(ret, c)
			}
		}
		return ret
	}
}

func Deck(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card
		for i := 0; i < n; i++ {
			ret = append(ret, cards...)
		}
		return ret
	}
}
