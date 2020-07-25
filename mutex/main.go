package main

import (
	"fmt"
	"sync"
	"time"
)

var count int
var lock sync.Mutex

func increment(num int) {
	lock.Lock()
	defer lock.Unlock()
	fmt.Printf("go-%d-increment-start\n", num)
	time.Sleep(100 * time.Millisecond)
	count++
	fmt.Printf("go-%d-increment-end : %d\n", num, count)
}

func decrement(num int) {
	lock.Lock()
	defer lock.Unlock()
	fmt.Printf("go-%d-decrement-start\n", num)
	time.Sleep(100 * time.Millisecond)
	count--
	fmt.Printf("go-%d-decrement-end : %d\n", num, count)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			increment(i)
		}(i)
	}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			decrement(i)
		}(i)
	}
	wg.Wait()
}
