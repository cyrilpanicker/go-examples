package main

import "fmt"

func sum(a float32, b float32) float32 {
	return a + b
}

func main() {
	integer1, integer2 := 1, 2
	fmt.Println(sum(float32(integer1), float32(integer2)))
}
