package main

import "fmt"

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func variadicTotal(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func goodbye() {
	fmt.Println("Goodbye!")
}

func main() {
	defer goodbye()

	xs := []float64{98,93,77,82,83}
	fmt.Println(average(xs))

	fmt.Println(variadicTotal(1,2,3,4,5))

	fmt.Println(factorial(3))
}