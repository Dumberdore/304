package game

import (
	"fmt"
)

func (G Game) DealRound() {
	fmt.Println("Dealing 4 cards to each player")

	players := G.Players
	for i := 0; i < 4; i++ {
		for _, p := range players {
			c := G.Deck.Draw()
			p.Cards = append(p.Cards, c)
		}
	}
}
