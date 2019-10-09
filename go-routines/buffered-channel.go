package main

import "fmt"

func bufferedChannel() {
	fmt.Println("-----------buffered channel starts-----------")
	channel := make(chan int, 3) //Sends block when the buffer is full. Receives block when the buffer is empty.
	channel <- 1
	channel <- 2
	channel <- 3
	// channel <- 4 //will block and exit because the buffer is full and the other side will never be ready to empty the buffer
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	// fmt.Println(<-channel) //will block and exit because the buffer is empty and the other side will never be ready to fill the buffer
	fmt.Println("-----------buffered channel ends-----------")
}
