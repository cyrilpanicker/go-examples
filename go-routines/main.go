package main

import (
	"fmt"
	"time"
)

func sendThroughChannel(value int, duration int, channel chan int) {
	time.Sleep(time.Duration(duration) * time.Millisecond)
	channel <- value
}

func returnNumbers(max int) []int {
	result := []int{}
	for index := 0; index < max; index++ {
		result = append(result, index)
	}
	return result
}

func receiveNumbers(max int, channel chan int) {
	for index := 0; index < max; index++ {
		channel <- index
	}
	close(channel)
}

func main() {
	channel := make(chan int)
	go sendThroughChannel(5, 100, channel)
	go sendThroughChannel(10, 200, channel)
	go sendThroughChannel(15, 0, channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	// blockingChannel()
	bufferedChannel()
	for _, value := range returnNumbers(5) {
		fmt.Printf("%v, ", value)
	}
	fmt.Println()
	channel1 := make(chan int, 5)
	go receiveNumbers(cap(channel1), channel1)
	for value := range channel1 { //will block until buffer is filled by sender
		fmt.Printf("%v, ", value)
	}
	fmt.Println()
	selectChannel()
}
