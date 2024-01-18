package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	card := "Ace of spades"
	new_card := newCard()
	fmt.Println(card, new_card)

}

func newCard() string {
	return " new Cards"
}
