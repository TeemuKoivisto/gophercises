package main

import (
	"fmt"
	. "github.com/TeemuKoivisto/gophercises/deck-of-cards/deck"
	"sort"
)

func createDeck() []Card {
	cards := make([]Card, 52)
	suit, suits := DIAMOND, 4
	for s := 0; s < suits; s++ {
		for i := 0; i < 13; i++ {
			cards[s*13+i] = Card{
				number: i + 1,
				suit:   suit,
			}
		}
		if suit == DIAMOND {
			suit = CLUBS
		} else if suit == CLUBS {
			suit = HEARTS
		} else if suit == HEARTS {
			suit = SPADES
		}
	}
	return cards
}

type ByAge []Card

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].number < a[j].number }

func main() {
	cards := createDeck()
	sort.Sort(ByAge(cards))
	fmt.Println(cards)
}
