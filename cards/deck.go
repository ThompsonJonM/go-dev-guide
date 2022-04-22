package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// Create a new type of "deck" which is a slice of strings
type deck []string

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

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

func (d deck) shuffle() {
	// Go always uses the same seed so we must generate a new seed source
	// then use the source for Rand
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (d deck) toString() string {
	// Do not need type conversion as deck extends type []string
	return strings.Join(d, ", ")
}

func createNewDeck() deck {
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

func newDeckFromFile(filename string) (deck, error) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return deck{}, err
	}

	splitStrings := strings.Split(string(bs), ", ")

	// Do not need type conversion as deck extends type []string
	return splitStrings, nil
}
