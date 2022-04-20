package main

func main() {
	cards := create()

	hand, _ := deal(cards, 3)
	hand.print()

	cards.saveToFile("my_cards")

	deck, _ := newDeckFromFile("my_cards")
	deck.print()
}
