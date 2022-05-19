package main

import "testing"

func TestCreateNewDeck(t *testing.T) {
	deck := createNewDeck()

	if len(deck) != 52 {
		t.Errorf("Expected deck to be size %d, got %d", 52, len(deck))
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected %s, got %s", "Ace of Spades", deck[0])
	}
}