package main

import "fmt"

func generateNumbers(num int, channel chan<- int) {
	defer close(channel)
	defer fmt.Println("generate-numbers-complete")
	for i := 0; i < num; i++ {
		fmt.Printf("generate-numbers-sending-%d\n", i)
		channel <- i
	}
}

func main() {
	channel := make(chan int, 100)
	go generateNumbers(10, channel)
	for num := range channel {
		fmt.Printf("main-received-%d\n", num)
	}
}
