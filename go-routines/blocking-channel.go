package main

import "fmt"

func blockingChannel() {
	channel := make(chan int) //sends and receives block until the other side is ready.
	channel <- 1              //will block and exit because the other side will never be ready.
	fmt.Println(<-channel)
}
