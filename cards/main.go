package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"

	// cards := newDeck()
	// hand, remainCards := deal(cards, 5)
	// fmt.Println("printing cards")
	// cards.print()
	// fmt.Println("***************")
	// hand.print()
	// fmt.Println("printing the remaing cards")
	// fmt.Println(hand.toString())

	// remainCards.print()
	// cards.saveToFile("save_cards")
	// cards := newDeckFromFile("save_cards")
	// cards.print()

	// cards := newDeck()
	// cards.shuffle()

	// cardsString := cards.toString()

	// fmt.Println(cardsString)

	// myNumbers := createNumOneToTen()
	// myEvens, myOdds := myNumbers.getEvenAndOdds()
	// printMyNums(myEvens)
	// printMyNums(myOdds)
	department := dept{deptId: 123, dName: "engineering"}
	eric := employee{name: "eric", age: 3000, dept: department}
	ericPtr := &eric
	ericPtr.updateName("eric cartman")

	fmt.Println(getGenericName(department))
	fmt.Println(getGenericName(eric))
	// mySlice := []int{1, 2, 3, 4, 5, 6}
	// udpateMySlice(mySlice)
	// fmt.Println(mySlice)

	//colors:=make(map[string]string) //creates empty map
	// var colors map[string]string //creates empty map
	countriesCaptial := map[string]string{
		"india":         "new delhi",
		"united states": "washingtondc",
		"pakistan":      "islamabad",
	}
	printMap(countriesCaptial)

}

func printMap(m map[string]string) {
	m["Sri Lanka"] = "Coimbture"
	for key, value := range m {
		fmt.Println(key, value)
	}
}

func udpateMySlice(mySlice []int) {
	mySlice[0] = 78
}
