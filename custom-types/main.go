package main

import "fmt"

type integer int

type point struct {
	x int
	y int
}

type points []struct {
	x int
	y int
}

type pointCollection map[string][]struct {
	x int
	y int
}

func main() {
	var a integer = 1
	b := integer(2)
	fmt.Println(a, b)
	c := point{x: 1, y: 2}
	fmt.Println(c)
	d := points{
		{x: 1, y: 2},
		{x: 3, y: 4},
	}
	fmt.Println(d)
	e := pointCollection{
		"key1": {{x: 1, y: 2}, {x: 3, y: 4}},
		"key2": {{x: 5, y: 6}, {x: 7, y: 8}, {x: 9, y: 10}},
	}
	fmt.Println(e)
}
