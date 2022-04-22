package main

import "log"

func main() {
	cards := createNewDeck()
	cards.shuffle()

	hand, _ := cards.deal(5)
	hand.print()

	err := cards.saveToFile("my_cards")
	if err != nil {
		log.Fatalf("%s", err)
	}

	_, err = newDeckFromFile("my_cards")
	if err != nil {
		log.Fatalf("%s", err)
	}
}
