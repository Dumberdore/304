package game

import (
	"fmt"
	"github.com/prakashdumbre-toast/304/Deck"
	"github.com/prakashdumbre-toast/304/player"
)

type Game struct {
	Deck    *Deck.Deck
	Players []*player.Player
	Hands   []Hand
	Risker  *player.Player
	Partner *player.Player
	Target  int
}

type Hand struct {
	Cards  []Deck.Card
	Winner *player.Player
}

func NewGame() {

	G := Game{}

	//	Create a Deck for the Game
	fmt.Println("Creating Deck")
	Deck := Deck.NewDeck()
	G.Deck = &Deck

	fmt.Println("Initiating Players")
	G.CreatePlayersAndAssignDefault()

	//	Deal 4 Cards to Each Player
	fmt.Println("Dealing 4 Cards to each Player")
	G.DealRound()

	// Round 1 of betting
	G.Round1Betting()

	//	Select Partner
	G.SelectPartner()
}

func (G Game) CreatePlayersAndAssignDefault() {
	//	Create 4 Players TODO: Move to init()
	fmt.Println("Creating 4 Players with P1 as Default")
	P1 := player.Player{Name: "Alpha"}
	P2 := player.Player{Name: "Bravo"}
	P3 := player.Player{Name: "Charlie"}
	P4 := player.Player{Name: "Delta"}

	G.Players = []*player.Player{&P1, &P2, &P3, &P4}

	// TODO : Create logic to keep track of defaults
	G.Risker = &P1
	G.Target = 180
}

func (G Game) Round1Betting() {
	//TODO
}

func (G Game) SelectPartner() {
	//TODO: Create logic to set partner
	partner := *G.Players[2]
	G.Partner = &partner
}
