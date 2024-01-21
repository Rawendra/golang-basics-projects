package main

import "fmt"

type myNums []int

func createNumOneToTen() myNums {
	count := 10
	n := myNums{}
	for i := 0; i < count; i++ {
		n = append(n, i)
	}
	return n
}

func (n myNums) getEvenAndOdds() (myNums, myNums) {
	evens := myNums{}
	odds := myNums{}
	for _, num := range n {
		if num%2 == 0 {
			evens = append(evens, num)
		} else {
			odds = append(odds, num)
		}
	}
	return evens, odds
}

func printMyNums(n myNums) {
	for _, num := range n {
		fmt.Println(num)
	}
}
