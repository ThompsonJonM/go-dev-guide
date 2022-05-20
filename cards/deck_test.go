package main

import (
	"os"
	"testing"
)

func TestCreateNewDeck(t *testing.T) {
	deck := createNewDeck()

	if len(deck) != 52 {
		t.Errorf("Expected deck to be size %d, got %d", 52, len(deck))
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected %s, got %s", "Ace of Spades", deck[0])
	}
}

func TestShuffleDeck(t *testing.T) {
	deck := createNewDeck()
	card := deck[0]

	deck.shuffle()

	if deck[0] == card {
		t.Errorf("Deck did not shuffle")
	}
}

func TestSaveToDeck(t *testing.T) {
	deck := createNewDeck()

	if err := deck.saveToFile("_decktesting"); err != nil {
		t.Errorf(err.Error())
	}

	if _, err := os.ReadFile("_decktesting"); err != nil {
		t.Errorf(err.Error())
	}

	if err := os.Remove("_decktesting"); err != nil {
		t.Errorf(err.Error())
	}
}

func TestNewDeckFromFile(t *testing.T) {
	deck := createNewDeck()

	if err := deck.saveToFile("_decktesting"); err != nil {
		t.Errorf(err.Error())
	}

	loadedDeck, err := newDeckFromFile("_decktesting")
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(loadedDeck) != 52 {
		t.Errorf("Expected %d, got %d", 52, len(loadedDeck))
	}

	if err := os.Remove("_decktesting"); err != nil {
		t.Errorf(err.Error())
	}
}