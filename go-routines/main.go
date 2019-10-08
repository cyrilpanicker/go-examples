package main

import (
	"fmt"
	"time"
)

func sendThroughChannel(value int, duration int, channel chan int) {
	time.Sleep(time.Duration(duration) * time.Millisecond)
	channel <- value
}

func main() {
	channel := make(chan int)
	go sendThroughChannel(5, 1000, channel)
	go sendThroughChannel(10, 2000, channel)
	go sendThroughChannel(15, 0, channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
