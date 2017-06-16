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
	fmt.Println("\n---------------------------")

	xs := filter([]int{1, 2, 3, 4}, func(n int) bool {
		return n > 1
	})
	fmt.Println(xs)

	fmt.Println("---------------------------")

	// anonymous self executed function
	func() {
		fmt.Println("anonymous function here!")
	}()

	// defer change func call order
	// right before the main exits the defers will be called
	defer world()
	hello()

}

func world() {
	fmt.Println("World")
}

func hello() {
	fmt.Print("Hello ")
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

func filter(numbers []int, callback func(int) bool) []int {
	xs := []int{}
	// or init as var xs []int
	for _, n := range numbers {
		if callback(n) {
			// add the new int var into xs array
			xs = append(xs, n)
		}
	}
	return xs
}
