package Deck

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Deck struct {
	RemainingCards int    `json:"remaining"`
	DeckId         string `json:"deck_id"`
}

type Card struct {
	Face string `json:"value"`
	Suit string `json:"suit"`
	Code string `json:"code"`
}

func NewDeck() Deck {
	GameDeck := Deck{}

	baseURL := "http://localhost:8080/api/"
	endpoint := "deck/new/shuffle/"
	param := "?cards=JS,9S,KS,QS,AS,0S,7S,8S,JD,9D,KD,QD,AD,0D,7D,8D,JC,9C,KC,QC,AC,0C,7C,8C,JH,9H,KH,QH,AH,0H,7H,8H"
	url := baseURL + endpoint + param

	DoCClient := http.Client{Timeout: time.Second * 2}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := DoCClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonErr := json.Unmarshal(body, &GameDeck)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return GameDeck
}

func (d *Deck) Draw(c int) []Card {
	if c == 0 {
		c = 1
	}
	baseURL := "http://localhost:8080/api/"
	endpoint := fmt.Sprintf("deck/%s/draw/?count=%d", d.DeckId, c)
	url := baseURL + endpoint

	DoCClient := http.Client{Timeout: time.Second * 2}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := DoCClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	type respDraw struct {
		Cards          []Card `json:"cards"`
		RemainingCards int    `json:"remaining"`
	}
	DrawnCard := respDraw{}

	jsonErr := json.Unmarshal(body, &DrawnCard)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	d.RemainingCards = DrawnCard.RemainingCards
	return DrawnCard.Cards
}
