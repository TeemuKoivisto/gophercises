package blackjack

import (
	"fmt"
)

func RunGame() GameState {
	game := New(GameOpts{})
	game.setPlayer(HUMAN, "Bob")
	game.setDealer(AI, "Janet")
	game.startGame()
	playerAI := newPlayerAI("Bob")
	dealerAI := newDealerAI("Janet")
	running := true
	for running {
		var input string
		switch game.state.Status {
		case PLAYER_TURN:
			input = playerAI.getAIMove(&game.state)
			switch input {
			case "h":
				game.handlePlayerHit()
			case "s":
				game.handlePlayerStand()
			}
		case DEALER_TURN:
			dealerMove := dealerAI.getAIMove(&game.state)
			if dealerMove == "h" {
				game.handleDealerHit()
			} else {
				game.endGame()
			}
		default:
			running = false
		}
	}
	return game.state
}

func Simulate() {
	fmt.Println("### BLACKJACK SIMULATOR ###")
	fmt.Println("Run how many simulations: ")
	var err error
	var input, wins, losses, draws int
	running := true
	for running {
		_, err = fmt.Scanf("%d", &input)
		if err == nil {
			running = false
		}
	}
	var result GameState
	for i := 0; i < input; i++ {
		result = RunGame()
		if result.Status == PLAYER_WON {
			wins += 1
		} else if result.Status == DEALER_WON {
			losses += 1
		} else {
			draws += 1
		}
	}
	fmt.Println("### RESULT ###")
	fmt.Println("Wins: ", wins)
	fmt.Println("Losses: ", losses)
	fmt.Println("Draws: ", draws)
	fmt.Println("%", float32(wins+1)/float32(wins+losses+1)*100)
}
