package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3, 8)

	// Fan out
	// distribute the sq work across two goroutines that both read from in
	c1 := sq(in)
	c2 := sq(in)

	// fan in
	// consume the merged output from multiple channels
	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
}

func gen(nums ...int) chan int {
	fmt.Printf("Type of nums: %T\n", nums)
	out := make(chan int)
	go func() {
		// _: index; n: value;
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...chan int) chan int {
	fmt.Printf("TYPE of cs: %T\n", cs) // just FYI
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func(ch chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
