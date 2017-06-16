package main

import "fmt"

func main() {
	// basic
	fmt.Println(greet("Fry", "Doe", 42))

	// variadic
	n := average(43, 56, 87, 12, 45, 57)
	fmt.Println(n)

	data := []float64{43, 56, 87, 72, 45, 47, 99}
	m := average(data...)
	fmt.Println(m)

	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Print(n, " ")
	})
	fmt.Println()
}

func greet(fname, lname string, lol int) (string, int) {
	return fmt.Sprint("Hey ", fname, " ", lname), lol
}

func average(sf ...float64) float64 {
	// sf -> array of float64 elements
	var total float64
	// foreach loop through sf
	for _, v := range sf {
		total += v
	}
	// length of sf array convert into float64
	return total / float64(len(sf))
}

func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}
