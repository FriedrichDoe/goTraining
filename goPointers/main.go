package main

import "fmt"

const metersToYards float64 = 1.09361

func main() {
	x := 42
	fmt.Println("x -", x)
	// prints where x variable is stored in memory
	fmt.Println("x memory adress -", &x)
	fmt.Printf("x memory adress as decimal - %d\n", &x)
	fmt.Println("-----------------------------------------")

	a := 43
	fmt.Println(a)  // 43
	fmt.Println(&a) // 0xc42000e248

	var b *int = &a // reference

	fmt.Println(b)  // 0xc42000e248
	fmt.Println(&b) // 0xc42000c030
	fmt.Println(*b) // dereference -> 43

	*b = 99
	fmt.Println(*b) // 99

	fmt.Println("-----------------------------------------")
	var meters float64
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters) // saves the value into the memory adress
	yards := meters * metersToYards
	fmt.Println(meters, "meters is", yards, "yards")
	fmt.Println("-----------------------------------------")

}
