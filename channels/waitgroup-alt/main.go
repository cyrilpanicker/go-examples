package main

import "fmt"

func greet(greeting string, done chan<- bool) {
	fmt.Println(greeting)
	done <- true
}

func main() {
	greetings := []string{"hi", "hello", "good day"}
	done := make(chan bool)
	for _, greeting := range greetings {
		go greet(greeting, done)
	}
	for range greetings {
		<-done
	}
}
