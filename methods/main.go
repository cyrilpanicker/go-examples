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

func scaleFunction(p *point, scale float64) {
	p.x = p.x * scale
	p.y = p.y * scale
}

func absFunction(p point) float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

//You can only declare a method with a receiver whose type is defined in the same package as the method.
//You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).
//In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.

func main() {

	point1 := point{1, 2}
	point2 := point{3, 4}
	fmt.Println(point1.abs(), point2.abs(), (&point1).abs(), (&point2).abs())
	fmt.Println(absFunction(point1), absFunction(point2))

	point3 := point{5, 6}
	fmt.Println(point3)
	point3.scale(2)
	fmt.Println(point3)
	(&point3).scale(0.5)
	fmt.Println(point3)
	scaleFunction(&point3, 2)
	fmt.Println(point3)

	integer1 := integer(-5)
	integer2 := integer(5)
	fmt.Println(integer1.abs(), integer2.abs())

}
