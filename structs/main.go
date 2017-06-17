package main

import "fmt"

// create own types
type foo int

type person struct {
	name string
	age  int
}

func (p person) fullName() string {
	return p.name + " Doe"
}

func main() {
	fmt.Println("Structs")

	var myAge foo
	myAge = 28
	fmt.Printf("%T %T %v \n", myAge, int(myAge), myAge)

	p1 := person{"Fry", 22}
	p2 := &person{"Pond", 54}

	fmt.Println(p1.fullName(), p1.age)
	fmt.Printf("%v %T \n", p2, p2)
}
