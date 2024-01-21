package main

import (
	"fmt"
)

func main() {
	rectangle := Reactangle{height: 40, width: 30}
	square := Square{side: 50}
	circle := Circle{radius: 7}
	fmt.Println("area of reactangle", getArea(rectangle))
	fmt.Println("area of square", getArea(square))
	fmt.Println("area of circle", getArea(circle))

}
