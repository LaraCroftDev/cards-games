package main

import "fmt"

const (
	playerCount = 4 // This is a 4 player game
)

func (d Deck) Deal() [][]card {
	// Shuffle the cards before dealing
	d.Shuffle()
	hands := make([][]card, playerCount)
	index := 0

	// First round: distributing 5 cards per palyer
	for i := 0; i < playerCount; i++ {
		hands[i] = append(hands[i], d.singleHand(index, 5)...)
		index += 5
	}

	// Continuing rounds: distributing 4 cards per palyer
	for k := 2; k > 0; k-- {
		for i := 0; i < playerCount; i++ {
			hands[i] = append(hands[i], d.singleHand(index, 4)...)
			index += 4
		}
	}
	// err := debugTrumpHands(hands)
	return hands
}

func (d Deck) singleHand(index, count int) []card {
	return d[index : index+count]
}

func debugTrumpHands(hands [][]card) error {
	unique := make(map[card]bool, cardsCount)
	for _, hand := range hands {
		for _, card := range hand {
			if !unique[card] {
				unique[card] = true
				continue
			}
		}
	}
	return nil
}

func printTrump(hands [][]card) {
	for i, d := range hands {
		fmt.Println("Player ", i+1, " cards: ")
		Deck(d).print()
	}
}
