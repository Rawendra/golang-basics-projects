package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

//create a new type of deck that is slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}
	cardfaces := []string{"Spades", "Hearts", "Diamond", "Club"}
	cardValues := []string{"Ace", "2", "3", "4"}

	for _, face := range cardfaces {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+face)
		}
	}

	return cards

}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// arguments first and then type
func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}
func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("ERRROR:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)

}

func (d deck) shuffle() {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	totalLength := len(d)

	for index := range d {
		randomIndex := r.Intn(totalLength - 1)
		d[index], d[randomIndex] = d[randomIndex], d[index]
	}
}
