package main

import "fmt"

func main() {

	// exercise 1
	var x [5]int
    x[4] = 100
	fmt.Println(x[4])

	// exercise 2
	fmt.Println(len(make([]int, 3, 9)))

	// exercise 3
	y := [6]string{"a","b","c","d","e","f"}
	fmt.Println(y[2:5])

	// exercise 4
	z := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}
	smallest := z[0]
	for _, value := range z {
		if value < smallest {
			smallest = value
		}
	}
	fmt.Println(smallest)
}