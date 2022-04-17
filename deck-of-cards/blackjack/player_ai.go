package blackjack

import "github.com/TeemuKoivisto/gophercises/deck-of-cards/deck"

type PlayerAI struct {
	name string
}

func newPlayerAI(name string) *PlayerAI {
	return &PlayerAI{
		name,
	}
}

func (ai *PlayerAI) getAIMove(state *GameState) string {
	pScore, dCard, soft := state.PlayerCards.Score(), state.DealerCards[0].Rank, Soft(state.PlayerCards...)
	if soft {
		if pScore < 18 {
			return "h"
		} else if pScore == 18 && dCard >= 9 {
			return "h"
		}
		return "s"
	} else if pScore < 12 {
		return "h"
	} else if pScore == 12 {
		if dCard == deck.Four || dCard == deck.Five || dCard == deck.Six {
			return "s"
		}
		return "h"
	} else if pScore == 13 || pScore == 14 {
		if dCard < 7 {
			return "s"
		}
		return "h"
	} else if pScore == 15 {
		if dCard < 7 || dCard == 10 {
			return "s"
		}
		return "h"
	} else if pScore == 16 {
		if dCard < 7 || dCard == deck.Nine || dCard == 10 || dCard == deck.Ace {
			return "s"
		}
		return "h"
	} else {
		return "s"
	}
}
