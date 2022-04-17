package blackjack

type PlayerAI struct {
	name string
}

func newPlayerAI(name string) *PlayerAI {
	return &PlayerAI{
		name,
	}
}

// func (ai *PlayerAI) getAIMove(state *GameState) string {
// 	pScore, dScore := state.PlayerCards.Score(), state.PlayerCards.Score()
// 	if dScore < 17 {
// 		return "h"
// 	} else {
// 		return "s"
// 	}
// }
