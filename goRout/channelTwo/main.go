package main

import "fmt"

// semaphore example
func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 10; i < 20; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		// <-done should be in a go routine!!!!
		// otherwise it can be stucked
		// deadlock
		<-done
		<-done
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}

}
