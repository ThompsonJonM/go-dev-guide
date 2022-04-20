package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new type of "deck" which is a slice of strings
type deck []string

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) saveToFile(filename string) error {
	if err := ioutil.WriteFile(filename, []byte(d.toString()), 0666); err != nil {
		return err
	}

	return nil
}

func (d deck) toString() string {
	// Do not need type conversion as deck extends type []string
	return strings.Join(d, ", ")
}

func create() deck {
	cards := deck{}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	cardSuits := []string{"Spades", "Clubs", "Hearts", "Diamonds"}

	for _, suit := range cardSuits {
		for _, val := range cardValues {
			card := fmt.Sprintf("%s of %s", val, suit)

			cards = append(cards, card)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func newDeckFromFile(filename string) (deck, error) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return deck{}, err
	}

	splitStrings := strings.Split(string(bs), ", ")
	d := deck(splitStrings)

	return d, nil
}
