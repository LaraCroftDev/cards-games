package main

import "fmt"

const fileName = "stored-cards.txt"

func main() {
	newDeck := NewDeck()
	shuffledDeck := Shuffle(newDeck)
	Printer(shuffledDeck)

	hands := Deal(newDeck)
	for i, hand := range hands {
		fmt.Println("Player: ", i+1, " Hand: ", hand)
	}

	FileSaver(shuffledDeck, fileName)
	FileSaver(hands, fileName)
}
