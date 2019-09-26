package main

import "fmt"

func readFile() {
	defer fmt.Println("closing file")
	fmt.Println("opening file")
	fmt.Println("reading file")
}

func deferSequence() {
	defer fmt.Println()
	for index := 0; index < 10; index++ {
		defer fmt.Printf("%d ", index)
	}
}

func main() {
	readFile()
	deferSequence()
}
