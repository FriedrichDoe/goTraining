package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrNorgateMath = errors.New("global defined error variable")

// create new error example
func main() {
	// first
	_, err := sqrt(-10)
	if err != nil {
		log.Println(err)
	}
	// second
	_, err2 := sqrt2(-10)
	if err2 != nil {
		log.Println(err2)
	}
	// third
	_, err3 := sqrt3(-30)
	if err3 != nil {
		log.Println(err3)
	}
}

// create error in if statement
func sqrt(f float64) (float64, error) {
	if f < 0 {
		// new error value
		return 0, errors.New("standart local error vriable")
	}

	// implementation
	return 42, nil
}

// use global error variable
func sqrt2(f float64) (float64, error) {
	if f < 0 {
		// new error value
		return 0, ErrNorgateMath
	}

	// implementation
	return 42, nil
}

// give the error additional variables for more information
func sqrt3(f float64) (float64, error) {
	if f < 0 {
		// new error value
		return 0, fmt.Errorf("Error with additional variable: %v", f)
	}

	// implementation
	return 42, nil
}
