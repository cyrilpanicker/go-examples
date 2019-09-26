package main

import "fmt"

func swap(a string, b string) (string, string) {
	return b, a
}

func main() {
	fmt.Println(swap("World", "Hello"))
}
