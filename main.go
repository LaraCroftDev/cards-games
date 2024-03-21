package main

import "fmt"

const fileName = "stored-cards.txt"

func main() {
	deck := NewDeck()
	deck.print()

	deck.Shuffle()
	deck.print()

	k := deck.Deal()
	fmt.Println("this is k: ", k)

}
