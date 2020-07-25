package main

import (
	"fmt"
	"sync"
)

func greet(greeting string) {
	fmt.Println(greeting)
}

func main() {
	var wg sync.WaitGroup
	greetings := []string{"hi", "hello", "good day"}
	for _, greeting := range greetings {
		wg.Add(1)
		go func(greeting string) {
			defer wg.Done()
			greet(greeting)
		}(greeting)
	}
	wg.Wait()
}
