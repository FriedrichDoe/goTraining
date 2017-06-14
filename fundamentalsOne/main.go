package main

import "fmt"

var g string
var h string = "protein"
var j, o string = "stored in j", "stored in o"

/**
  variables
*/
func main() {

	a := 10
	b := "golang"
	c := 4.17
	d := true

	fmt.Printf("%v \n", a)
	fmt.Printf("%T \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%T \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%T \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%T \n", d)

	fmt.Println("------------------------")
	fmt.Println("---undefined variables--")
	fmt.Println("------------------------")

	var q int
	var w string
	var r float64
	var t bool

	fmt.Printf("%v \n", q)
	fmt.Printf("%v \n", w)
	fmt.Printf("%v \n", r)
	fmt.Printf("%v \n", t)
	fmt.Printf("%v \n", nil)

}
