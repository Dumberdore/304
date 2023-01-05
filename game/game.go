package game

import (
	"fmt"
	"github.com/prakashdumbre-toast/304/Deck"
	"github.com/prakashdumbre-toast/304/player"
)

type Game struct {
	Deck      *Deck.Deck
	Players   []player.Player
	Tricks    []Trick
	Risker    *player.Player
	Partner   *player.Player
	GameCards []Deck.Card
	Target    int
}

var G = Game{
	Players: createPlayers(),
}

func NewGame() {

	// Assign Defaults
	G.Risker = &G.Players[0]
	G.Partner = &G.Players[2]

	//	Create a Deck for the Game
	fmt.Println("Creating Deck")
	D := Deck.NewDeck()
	G.Deck = &D

	//	Deal 4 Cards to Each Player
	G.GameCards = append(G.GameCards, G.Deck.Draw(16)...)
	G.DealRound1()

	// Round 1 of betting
	G.Round1Betting()

	//	Select Partner
	G.SelectPartner()

	//	Deal remaining 4 cards to each player
	G.GameCards = append(G.GameCards, G.Deck.Draw(16)...)
	G.DealRound2()

	fmt.Println("After Dealing Round 2, players are as follow-")
	for _, p := range G.Players {
		fmt.Println("")
		fmt.Printf("Name : %s", p.Name)
		fmt.Printf("Cards : %v", p.HoldingCards)
		fmt.Println("--------------------------------")
	}

	fmt.Println("GameCards : ", G.GameCards)
	//	Collect Tricks
	G.CollectTricks()

}

func createPlayers() []player.Player {
	//	Create 4 Players TODO: Move to init()
	fmt.Println("Creating 4 Players with P1 as Default")
	Players := []player.Player{
		{
			Name:         "Alpha",
			HoldingCards: []*Deck.Card{},
		},
		{
			Name:         "Bravo",
			HoldingCards: []*Deck.Card{},
		},
		{
			Name:         "Charlie",
			HoldingCards: []*Deck.Card{},
		},
		{
			Name:         "Delta",
			HoldingCards: []*Deck.Card{},
		},
	}

	return Players
}

func (G *Game) Round1Betting() {
	//TODO
}

func (G *Game) SelectPartner() {
	//TODO: Create logic to set partner
}
