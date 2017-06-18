package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int64

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

		atomic.AddInt64(&counter, 1)

		fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}
