package main

import "fmt"
import "github.com/FriedrichDoe/goTraining/stringutil"

func main() {
	fmt.Println("Hello World!")

	fmt.Println(stringutil.MyName)

	// %d -> dezimal -> 42
	// %b -> binary -> 101010
	// TODO -> format variables:
	// https://godoc.org/fmt
	fmt.Printf("%d - %b - %x \n", 42, 42, 42)

	// \t   -> tabulator
	fmt.Printf("%d \t %b \t %#x \n", 42, 42, 42)

	// TODO -> for loop syntax
	// https://golang.org/doc/effective_go.html#for
	for index := 0; index < 200; index++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", index, index, index, index)
	}
}
