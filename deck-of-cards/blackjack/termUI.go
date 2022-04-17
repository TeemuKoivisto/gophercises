package blackjack

import (
	"fmt"
)

func Play() {
	game := New()
	game.setPlayer(HUMAN, "Bob")
	game.setDealer(AI, "Janet")
	game.startGame()
	fmt.Println("### BLACKJACK ###")
	running := true
	for running {
		var input string
		switch game.state.Status {
		case PLAYER_TURN:
			PrintStatus(game)
			input = PromptPlayerAction()
		case DEALER_TURN:
			// get ai move
			game.endGame()
		default:
			running = false
		}
		switch input {
		case "h":
			game.handlePlayerHit()
		case "s":
			game.handlePlayerStand()
		}
	}
	PrintGameEnd(game)
}

func PrintStatus(g *Game) {
	player, dealer := g.state.PlayerCards, g.state.DealerCards
	fmt.Println("Player:", player)
	fmt.Println("Dealer:", dealer.DealerString())
}

func PromptPlayerAction() string {
	var input string
	fmt.Println("Player turn: (h)it, (s)tand")
	fmt.Scanf("%s\n", &input)
	return input
}

func PrintGameEnd(g *Game) {
	pScore, dScore := g.state.PlayerCards.Score(), g.state.DealerCards.Score()
	fmt.Println("### GAME ENDED ###")
	fmt.Println("Player:", g.state.PlayerCards, "\nScore:", pScore)
	fmt.Println("Dealer:", g.state.DealerCards, "\nScore:", dScore)
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
