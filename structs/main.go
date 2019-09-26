package main

import "fmt"

type point struct {
	x int
	y int
}

func main() {

	point1 := point{3, 4}
	fmt.Println(point1, point1.x, point1.y)

	point1Pointer := &point1
	fmt.Println((*point1Pointer).x, point1Pointer.x, (*point1Pointer).x == point1.x)

	fmt.Println(point{1, 2}, point{y: 1}, point{})

	structLiteral := struct {
		a bool
		b bool
	}{true, false}
	fmt.Println(structLiteral)

}
