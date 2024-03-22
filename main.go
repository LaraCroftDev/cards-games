package main

const fileName = "stored-cards.txt"

func main() {
	deck := NewDeck()
	deck.print()

	deck.Shuffle()
	deck.print()

	k := deck.Deal()
	SaveToFile(k, fileName)
	printTrump(k)

	ReadDeckFromFile(fileName)
}
