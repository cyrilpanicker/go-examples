package main

import "fmt"

func compute(function func(int, int) int, a int, b int) int {
	return function(a, b)
}

func add(x int, y int) int {
	return x + y
}

func subtract(x int, y int) int {
	return x - y
}

func adder(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

func main() {
	fmt.Println(compute(add, 1, 2))
	fmt.Println(compute(subtract, 4, 3))
	fmt.Println(adder(1)(2))
	fmt.Println(adder(5)(0))
}
