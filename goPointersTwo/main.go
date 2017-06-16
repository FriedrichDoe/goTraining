package main

import "fmt"

// wrong function
func failZero(x int) {
	fmt.Println("fail Zero address", &x)
	x = 0
}

// right function
func zero(x *int) {
	*x = 0
}

func main() {
	x := 5
	fmt.Println("original x address", &x) // address in main
	failZero(x)
	fmt.Println(x) // x is still 5

	fmt.Println("---------------------------")
	zero(&x) // give address of x
	fmt.Println(x)

}
