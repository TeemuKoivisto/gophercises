package main

import (
	"fmt"
	"strings"

	"github.com/TeemuKoivisto/gophercises/deck-of-cards/deck"
)

type Hand []deck.Card

type BlackJackStatus uint8

const (
	PLAYER_TURN BlackJackStatus = iota
	DEALER_TURN
	PLAYER_WON
	DEALER_WON
	DRAW
)

type BlackJackState struct {
	Status      BlackJackStatus
	PlayerCards Hand
	DealerCards Hand
}

type BlackJackGame struct {
	State BlackJackState
}

// func New() {
// 	return BlackJackGame{
// 		State: BlackJackState{
// 			Status: PLAYER_TURN,
// 		}
// 	}
// }

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

func (h Hand) DealerString() string {
	return h[0].String() + ", **HIDDEN**"
}

func (h Hand) Score() int {
	minScore := h.MinScore()
	if minScore > 11 {
		return minScore
	}
	for _, c := range h {
		if c.Rank == deck.Ace {
			return minScore + 10
		}
	}
	return minScore
}

func (h Hand) MinScore() int {
	score := 0
	for _, c := range h {
		score += min(int(c.Rank), 10)
	}
	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// func (b *BlackJackGame) HandleTurn() {
// 	switch (b.State.Status) {
// 	case PLAYER_TURN: {

// 	}
// 	}
// }

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

func BlackJack() {
	cards := deck.New(deck.Deck(3), deck.Shuffle)
	var card deck.Card
	var player, dealer Hand
	for i := 0; i < 2; i++ {
		for _, hand := range []*Hand{&player, &dealer} {
			card, cards = draw(cards)
			*hand = append(*hand, card)
		}
	}
	var input string
	fmt.Println("### BLACKJACK ###")
	for input != "s" {
		fmt.Println("Player:", player)
		fmt.Println("Dealer:", dealer.DealerString())
		fmt.Println("Player turn: (h)it, (s)tand")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			{
				card, cards = draw(cards)
				player = append(player, card)
			}
		}
	}
	pScore, dScore := player.Score(), dealer.Score()
	fmt.Println("### GAME ENDED ###")
	fmt.Println("Player:", player, "\nScore:", pScore)
	fmt.Println("Dealer:", dealer, "\nScore:", dScore)
	switch {
	case pScore > 21:
		fmt.Println("You lost. You busted")
	case dScore > 21:
		fmt.Println("You won. Dealer busted")
	case pScore > dScore:
		fmt.Println("You won.")
	case dScore > pScore:
		fmt.Println("Dealer won")
	case pScore == dScore:
		fmt.Println("Draw")
	}
}
