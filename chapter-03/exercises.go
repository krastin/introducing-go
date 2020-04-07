package main

import "fmt"

func main() {
	// exercise 1
	name := "John Smith"
	var age int
	age = 35

	//exercise 2
	x := 5
	x += 1
	fmt.Println(x) // 6

	//exercise 3
	// Scope of variables is their visibility according to the layers of functions they are defined in

	//exercise 4
	// const variables cannot be changed once assigned

}

//exercise5
func fIntoC(degreesF float64) float64 {
	return (degreesF - 32) * 5/9
}

//exercise6
func fIntoM(feet float64) float64 {
	return feet * 0.3048
}