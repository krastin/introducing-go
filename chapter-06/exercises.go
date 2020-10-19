package main

import "fmt"

// exercise 2
func halfEven(x int) (int, bool) {
	isEven := func(x int) bool {
		if (x % 2.0) > 0 {
			return false
		}
		return true
	}
	return x / 2, isEven(x / 2)
}

//exercise 3
func greatestNumber(args ...int) int {
	greatest := args[0]
	for _, number := range args {
		if number > greatest {
			greatest = number
		}
	}
	return greatest
}

//exercise 4
func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

//exercise 5
func fibb(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	//fmt.Println("returning fibb(", n-1, ") + fibb(", n-2, ")")
	return fibb(n-1) + fibb(n-2)
}

//exercise 11
func swap(x *int, y *int) {
	z := *y
	*y = *x
	*x = z
}

func main()  {
	fmt.Println(halfEven(8))
	fmt.Println(halfEven(10))

	fmt.Println(greatestNumber(1,2,3,100,4,5))

	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())

	fmt.Println(fibb(6))

	x := 5
	y := 10
	fmt.Println("x: ", x, "\t y: ", y)
	swap(&x, &y)
	fmt.Println("x: ", x, "\t y: ", y)
}