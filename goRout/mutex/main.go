package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex

// go run -race main.go
// check how many race conditions in the code
func main() {
	wg.Add(2)
	go incrementor("foo:")
	go incrementor("bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)

		mutex.Lock() // lock the variable
		counter++
		mutex.Unlock() // unlock the variable

		fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}
