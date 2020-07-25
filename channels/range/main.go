package main

import "fmt"

func generateIntegers(num int, channel chan<- int) {
	defer close(channel)
	for i := 0; i < num; i++ {
		channel <- i
	}
}

func main() {
	channel := make(chan int)
	go generateIntegers(10, channel)
	for num := range channel {
		fmt.Println(num)
	}
}
