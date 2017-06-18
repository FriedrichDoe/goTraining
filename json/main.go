package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

/*
   Person comment
*/
type Person struct {
	First string
	Last  string `json:"-"`            // invisible
	Age   int    `json:"wisdom score"` // rename
}

func main() {
	// test struct to marshal into json
	p1 := Person{"James", "Bond", 20}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Println(string(bs))

	// Unmarshal
	var p2 Person // empty struct with zero value
	fmt.Println("firstname:", p2.First, " // age:", p2.Age, " // ")

	bss := []byte(`{"First":"Winston","wisdom score":32}`)
	json.Unmarshal(bss, &p2)

	fmt.Println("--------------------")
	fmt.Println(p2.First)
	fmt.Println(p2.Last)
	fmt.Println(p2.Age)
	fmt.Printf("%T \n", p2)

	fmt.Println("--------------------")

	// encoding
	json.NewEncoder(os.Stdout).Encode(p2)
	fmt.Println("--------------------")

	// decoding
	var p3 Person
	rdr := strings.NewReader(`{"First":"Winston","wisdom score":32}`)
	json.NewDecoder(rdr).Decode(&p3)

	fmt.Println("p3:", p3.First, p3.Age)
}
