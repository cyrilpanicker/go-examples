package main

import (
	"fmt"
	"runtime"
	"time"
)

func isEven(x int) bool {
	if x%2 == 0 {
		return true
	}
	return false
}

func square(x int, limit int) int {
	if square := x * x; square < limit {
		return square
	}
	return limit
}

func os() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s\n", os)
	}
}

func greet() {
	time := time.Now()
	switch {
	case time.Hour() < 12:
		fmt.Println("Good Morning")
	case time.Hour() < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Evening")
	}
}

func main() {
	fmt.Println(isEven(5), isEven(10))
	fmt.Println(square(3, 10))
	fmt.Println(square(4, 10))
	os()
	greet()
}
