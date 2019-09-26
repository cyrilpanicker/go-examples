package main

import "fmt"

func a() int {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	return sum
}

func b() int {
	sum := 0
	i := 0
	for i < 10 {
		sum += i
		i++
	}
	return sum
}

func main() {
	fmt.Println(a())
	fmt.Println(b())
}
