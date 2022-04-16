package main

import (
	"fmt"
	"github.com/TeemuKoivisto/gophercises/deck-of-cards/deck"
)

func main() {
	cards := deck.New()
	cards = deck.Shuffle(cards)
	fmt.Println(cards)
}
