package blackjack

type DealerAI struct {
	name string
}

func newDealerAI(name string) *DealerAI {
	return &DealerAI{
		name,
	}
}

func (ai *DealerAI) getAIMove(state *GameState) string {
	dScore := state.DealerCards.Score()
	if dScore < 17 {
		return "h"
	} else {
		return "s"
	}
}
