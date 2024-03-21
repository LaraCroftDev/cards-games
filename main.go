package main

const fileName = "stored-cards.txt"

func main() {
	deck := NewDeck()
	deck.print()

	deck.Shuffle()
	deck.print()
}
