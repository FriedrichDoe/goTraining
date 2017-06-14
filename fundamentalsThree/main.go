package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// program to get an website and print its sourcecode
func main() {
	// get website / _ blank identifier for return type
	res, _ := http.Get("https://www.google.de")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	// %s print website as string
	fmt.Printf("%s \n", page)
}
