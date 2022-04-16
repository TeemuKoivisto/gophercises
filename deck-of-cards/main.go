package main

import (
	"fmt"
	. "github.com/TeemuKoivisto/gophercises/deck-of-cards/deck"
	"sort"
)

func createDeck() []Card {
	cards := make([]Card, 52)
	suit, suits := Diamond, 4
	for s := 0; s < suits; s++ {
		for i := 0; i < 13; i++ {
			cards[s*13+i] = Card{
				Rank: CardRank(i + 1),
				Suit: suit,
			}
		}
		if suit == Diamond {
			suit = Club
		} else if suit == Club {
			suit = Heart
		} else if suit == Heart {
			suit = Spade
		}
	}
	return cards
}

type ByAge []Card

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Rank < a[j].Rank }

func main() {
	cards := createDeck()
	sort.Sort(ByAge(cards))
	fmt.Println(cards)
}
