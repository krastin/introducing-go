package main

import "fmt"

func main() {
	defer func() {
		str := recover()
		fmt.Println("Panicking: ", str)
	}()
	panic("PANIC")
}