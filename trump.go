package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var ranks = [rankCount]string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}
var suits = [suitCount]string{"Heart", "Diamond", "Club", "Spade"}

const (
	suitCount   = 4
	rankCount   = 13
	cardsCount  = 52
	playerCount = 4 // Trump is a 4 player game
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

func (d Deck) Deal() []Deck {
	// Shuffle the cards before dealing
	d.Shuffle()
	hands := make([]Deck, playerCount)
	index := 0

	// First round: distributing 5 cards per player
	for i := 0; i < playerCount; i++ {
		hands[i] = append(hands[i], d.singleHand(index, 5)...)
		index += 5
	}
	// The two continuing rounds: distributing 4 cards per player
	for k := 2; k > 0; k-- {
		for i := 0; i < playerCount; i++ {
			hands[i] = append(hands[i], d.singleHand(index, 4)...)
			index += 4
		}
	}
	// err := debugTrumpHands(hands)
	return hands
}

// SaveToFile func can be used to write Deck and []Deck types of data into a text file
func SaveToFile(data interface{}, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Using a buffer to enhance the performance
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(toString(data))
	if err != nil {
		return err
	}
	writer.Flush()
	return nil
}

func ReadDeckFromFile(fileName string) error {
	b, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	read := string(b)
	fmt.Println("File", fileName, "content:\n", read)
	return nil
}

func (d Deck) print() {
	for _, card := range d {
		fmt.Println(card.rank, card.suit)
	}
}

func printTrump(hands []Deck) {
	for i, d := range hands {
		fmt.Println("Player ", i+1, " cards: ")
		d.print()
	}
}

func debugTrumpHands(hands []Deck) error {
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

func (d Deck) singleHand(index, count int) []card {
	return d[index : index+count]
}

func toString(data interface{}) string {
	switch d := data.(type) {
	case Deck:
		return joinCards(d)
	case []Deck:
		return joinHands(d)
	default:
		// log some error potentially
		return ""
	}
}

func joinCards(deck Deck) string {
	cards := make([]string, len(deck))
	for i, card := range deck {
		cards[i] = card.rank + card.suit
	}
	return strings.Join(cards, ", ")
}

func joinHands(hands []Deck) string {
	handStrings := make([]string, len(hands))
	for i, hand := range hands {
		handStrings[i] = joinCards(hand)
	}
	return strings.Join(handStrings, "\n")
}
