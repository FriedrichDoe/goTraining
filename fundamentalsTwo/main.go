package main

import "fmt"

var x int = 42

func main() {
	y := 69

	// anonymous function
	increment := func() int {
		y++
		return y
	}

	fmt.Println(increment())
	fmt.Println(increment())

	// other anonymous function
	increment2 := wrapper()

	fmt.Println(increment2())
	fmt.Println(increment2())

	fmt.Println(x)
	foo()

	max := max(55)
	fmt.Println(max)
}

func foo() {
	fmt.Println(x)
}

func max(x int) int {
	return 42 + x
}

// info: bb value will be saved after first call
// 1. run -> 4
// 2. run -> 16
func wrapper() func() int {
	bb := 2
	return func() int {
		bb *= bb
		return bb
	}
}
