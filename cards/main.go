package main

func main() {
	// cards := newDeck()
	// cards.saveToFile("My cards")

	cards := newDeckFromFile("My cards")
	cards.print()
}
