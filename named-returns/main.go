package main

import "fmt"

func sum(a int, b int) (sum int) {
	sum = a + b
	return
}

func main() {
	fmt.Println(sum(1, 2))
}
