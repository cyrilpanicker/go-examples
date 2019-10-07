package main

import (
	"fmt"
	"math"
)

type integer int

func (x integer) abs() integer {
	if x < 0 {
		return -x
	}
	return x
}

type point struct {
	x float64
	y float64
}

func (p point) abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func (p *point) scale(scale float64) {
	p.x = p.x * scale
	p.y = p.y * scale
}

//You can only declare a method with a receiver whose type is defined in the same package as the method.
//You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).

func main() {

	point1 := point{x: 1, y: 2}
	point2 := point{x: 3, y: 4}
	fmt.Println(point1.abs(), point2.abs())

	point3 := point{x: 5, y: 6}
	fmt.Println(point3)
	point3.scale(2)
	fmt.Println(point3)

	integer1 := integer(-5)
	integer2 := integer(5)
	fmt.Println(integer1.abs(), integer2.abs())

}
