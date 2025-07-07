package main

import "fmt"

func main() {
	// fmt.Println("hello goffers")
	// strings
	var nameOne string = "bode"
	var school = "Futa"
	var nameThree string
	fmt.Println(nameOne, school, nameThree)

	//Reassigning the variables
	nameOne = "Gboyega"
	nameThree = "Tayo"

	fmt.Println(nameOne, nameThree)

	//another way to declare a variable
	nameFive := "Oluwasegun"
	fmt.Println(nameFive)

	//Int
	var numberOne int = 20
	var numberTwo = 40
	numberThree := 60

	fmt.Println(numberOne, numberTwo, numberThree)

	//Floats
	var scoreOne float64 = 55.10
	scoreTwo := 99.60

	fmt.Println(scoreOne, scoreTwo)
}
