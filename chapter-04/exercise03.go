package main

import "fmt"

func main() {
	for i:=1; i<101; i++ {
		fizzbuzz := false
		if i % 3 == 0 {
			fmt.Print("Fizz")
			fizzbuzz = true
		}
		if i % 5 == 0 {
			fmt.Print("Buzz")
			fizzbuzz = true
		}
		if fizzbuzz {
			fmt.Println()
			continue
		}
		fmt.Println(i)
	}
}