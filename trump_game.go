package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// const (
// 	// 	suitNum    = 4
// 	// 	typeNum    = 13
// 	// 	cardsNum   = 52
// 	playersNum = 4 // This is a 4 player card game
// )

// // type Card struct {
// // 	Type string
// // 	Suit string
// // }

// // type Deck [cardsNum]Card

// // func NewDeck() Deck {
// // 	types := [typeNum]string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}
// // 	suits := [suitNum]string{"Heart", "Diamond", "Club", "Spade"}
// // 	var deck Deck
// // 	n := 0
// // 	m := 0
// // 	for i, card := range deck {
// // 		if m == typeNum {
// // 			n++
// // 			m = 0
// // 		}
// // 		card.Suit = suits[n]
// // 		card.Type = types[m]
// // 		m++

// // 		deck[i] = card
// // 	}
// // 	return deck
// // }

// func Printer(deck Deck) {
// 	fmt.Println(deck)
// }

// // func Shuffle(deck Deck) Deck {
// // 	for i := 0; i < cardsNum; i++ {
// // 		randIndex := rand.Intn(cardsNum)
// // 		// Swap current card with random Index card
// // 		deck[i], deck[randIndex] = deck[randIndex], deck[i]
// // 	}
// // 	return deck
// // }

// func Deal(deck Deck) [][]Card {
// 	// Shuffle the cards before dealing
// 	shuffledDeck := Shuffle(deck)
// 	hands := make([][]Card, suitNum)
// 	j := 0

// 	// First round: distributing 5 cards per palyer
// 	for i := 0; i < playersNum; i++ {
// 		hands[i] = shuffledDeck[j : j+5]
// 		j += 5
// 	}

// 	p := 0
// 	for k := j; k < cardsNum; k += 4 {
// 		if p == playersNum {
// 			p = 0
// 		}
// 		// Continuing rounds: distributing 4 cards per palyer
// 		hands[p] = append(hands[p], shuffledDeck[k:k+4]...)
// 		// Moving on to next player
// 		p++
// 	}
// 	return hands
// }

// func FileSaver(data interface{}, fileName string) error {

// 	file, err := os.Create(fileName)
// 	if err != nil {
// 		fmt.Println("Couldn't create a new file: ", err.Error())
// 		return err
// 	}

// 	defer file.Close()

// 	_, err = file.WriteString(stringConverter(data))
// 	if err != nil {
// 		fmt.Println("Couldn't write into file: ", err.Error())
// 		return err
// 	}
// 	return nil

// }

// func stringConverter(data interface{}) string {
// 	res := ""
// 	if cards, ok := data.([]Card); ok {
// 		for _, card := range cards {
// 			res = res + card.Type + card.Suit + ", "
// 		}
// 		return res
// 	}

// 	if deck, ok := data.(Deck); ok {
// 		for i := 0; i < len(deck); i++ {
// 			res = res + deck[i].Type + deck[i].Suit + ", "
// 		}
// 		return res
// 	}

// 	hands := data.([][]Card)
// 	for i, cards := range hands {
// 		res = res + "player " + strconv.Itoa(i+1) + " cards: " + stringConverter(cards) + "\n"
// 	}
// 	return res
// }
