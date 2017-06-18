package main

import (
	"fmt"
	"time"
)

func main() {
	unBufChannel()
	fmt.Print("\n")
	rangeCh()
}

func rangeCh() {
	r := make(chan int)

	go func() {
		for i := 10; i < 20; i++ {
			r <- i // code stops here until c will me used
		}
		close(r)
	}()

	for n := range r {
		fmt.Print(n, " ")
	}
	fmt.Print("\n")
}

func unBufChannel() {
	// unbuffered channel
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i // code stops here until c will me used
		}
	}()

	go func() {
		for {
			fmt.Print(<-c, " ") // c will be used here
		}
	}()

	time.Sleep(time.Second / 2)
}
