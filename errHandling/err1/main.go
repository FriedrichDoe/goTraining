package main

import (
	"fmt"
	"log"
	"os"
)

// logs into file
func init() {
	nf, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nf)
}

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {

		// different ways to output errors
		// fmt.Println("Error happened here", err)

		// todds favorite
		log.Println("err happended", err)
		// log.Fatalln(err)
		// panic(err)
	}
}
