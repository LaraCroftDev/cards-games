package main

const fileName = "stored-cards.txt"

func main() {
	deck := NewDeck()
	deck.print()
	err := SaveToFile(deck, fileName)
	if err != nil {
	}

	deck.Shuffle()
	k := deck.Deal()
	printTrump(k)
	err = SaveToFile(k, fileName)
	if err != nil {
	}
}
