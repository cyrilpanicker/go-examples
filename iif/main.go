package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("parallel test")
	}()
	func() {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("test")
	}()
}
