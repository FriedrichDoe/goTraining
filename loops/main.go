package main

import "fmt"

// https://golang.org/doc/effective_go.html#for
func main() {

	fmt.Println("--------------------------------------------")

	foo := 'a'
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)
	fmt.Printf("%T \n", rune(foo))

	// interessting stuff
	fmt.Println([]byte("Hello World"))

	fmt.Println("--------------------------------------------")

	// standart for loop
	for i := 0; i < 20; i++ {
		fmt.Println(i)
	}

	// while true loop
	for {
		fmt.Println("true")
		break
	}

	i := 0
	// for with condition
	for i < 10 {
		fmt.Println(i)
		i++
	}

	x := 0

	for {
		x++
		if x%2 == 0 {
			// overjump the rest of code below and start new loop
			continue
		}
		fmt.Println("continue loop: ", x)
		if x >= 20 {
			break
		}
	}
	fmt.Println("--------------------------------------------")
	for i := 5000; i <= 5100; i++ {
		fmt.Println(i, "-", string(i), "-", []byte(string(i)))
	}
	fmt.Println("--------------------------------------------")

	switch "Marcus" {
	case "Tim", "Artha":
		fmt.Println("Wassup Tim")
	case "Marcus":
		fmt.Println("Wassup Marcus")
		fallthrough
	case "Thomas":
		fmt.Println("Wassup Thomas")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	default:
		fmt.Println("have you no friends?")
	}

	var g interface{}
	g = 42

	switch g.(type) {
	case int:
		fmt.Println("g is int")
	case string:
		fmt.Println("g is string")
	case float32:
		fmt.Println("g is float32")
	}

}
