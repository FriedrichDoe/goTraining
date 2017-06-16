package main

import "fmt"

const (
	A = iota // 0
	B = iota // 1
	C = iota // 2
)

const (
	_ = iota      // 0
	D = iota * 10 // 1 * 10
	E = iota * 10 // 2 * 10
	F = iota * 10 // 3 * 10
)

const (
	_   = iota             // 0
	KB  = 1 << (iota * 10) // 1 << (1 * 10)
	MB  = 1 << (iota * 10) // 1 << (2 * 10)
	MTB = 4 << (iota * 10) // 1 << (3 * 10)
	// Constants Video
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)

	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", MTB)
	fmt.Printf("%d\n", MTB)
}
