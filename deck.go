package main

import (
	"fmt"
	"math/rand"
)

const (
	suitNum  = 4
	typeNum  = 13
	cardsNum = 52
)

type Card struct {
	Type string
	Suit string
}
type Deck []Card

func NewDeck() Deck {
	types := [typeNum]string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}
	suits := [suitNum]string{"Heart", "Diamond", "Club", "Spade"}

	deck := make(Deck, cardsNum)
	n := 0
	m := 0

	for i := 0; i < cardsNum; i++ {
		if m == typeNum {
			n++
			m = 0
		}
		deck[i].Suit = suits[n]
		deck[i].Type = types[m]
		m++
	}
	return deck
}

func (d Deck) Shuffle() {
	for i := 0; i < cardsNum; i++ {
		randIndex := rand.Intn(cardsNum)
		// Swap current card with random Index card
		d[i], d[randIndex] = d[randIndex], d[i]
	}
}

func (d Deck) Printer() {
	for _, card := range d {
		fmt.Println(card.Type, card.Suit)
	}
}
