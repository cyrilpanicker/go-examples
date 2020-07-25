package main

import "fmt"

func greet(channel chan int) {
	fmt.Println("hello")
	close(channel)
}

func main() {
	channel := make(chan int)
	go greet(channel)
	<-channel
}
