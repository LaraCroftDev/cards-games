package main

const fileName = "stored-cards.txt"

func main() {
	deck := NewDeck()
	deck.Shuffle()
	deck.Printer()

	// shuffledDeck := Shuffle(newDeck)

	// Printer(shuffledDeck)

	// hands := Deal(newDeck)
	// // for i, hand := range hands {
	// // 	fmt.Println("Player: ", i+1, " Hand: ", hand)
	// // }

	// FileSaver(shuffledDeck, fileName)
	// FileSaver(hands, fileName)
}
