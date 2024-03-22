package main

import (
	"fmt"
	"math/rand"
)

var ranks = [rankCount]string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}
var suits = [suitCount]string{"Heart", "Diamond", "Club", "Spade"}

const (
	suitCount  = 4
	rankCount  = 13
	cardsCount = 52
)

type card struct {
	rank string
	suit string
}

type Deck []card

func NewDeck() Deck {
	deck := make(Deck, cardsCount)
	deck = nil
	for _, s := range suits {
		for _, r := range ranks {
			deck = append(deck, card{rank: r, suit: s})
		}
	}
	return deck
}

func (d Deck) Shuffle() {
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}

func (d Deck) print() {
	for _, card := range d {
		fmt.Println(card.rank, card.suit)
	}
}
