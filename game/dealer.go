package game

import (
	"fmt"
)

func (G *Game) DealRound1() {
	fmt.Println("Dealing 4 cards to each player")

	println("Current Game Cards :")
	for i, c := range G.GameCards {
		fmt.Printf("Give Card %s to player index %s\n", c.Code, G.Players[i%4].Name)
		G.Players[i%4].HoldingCards = append(G.Players[i%4].HoldingCards, &c)
	}

	for _, p := range G.Players {
		fmt.Printf("\n[%s] : ", p.Name)
		for _, c := range p.HoldingCards {
			fmt.Printf("%s ", c.Code)
		}
	}

	fmt.Println()
}

func (G *Game) DealRound2() {

}

func (G *Game) CollectTricks() {
	//	Collect 8 Tricks
	//TODO
}
