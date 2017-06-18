package main

import "fmt"

type Square struct {
	side float64
}

func (z Square) area() float64 {
	return z.side * z.side
}

type Shape interface {
	area() float64
}

func main() {
	s := Square{13}
	info(s)
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}
