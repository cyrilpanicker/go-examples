package main

import (
	"fmt"
	"math"
)

type abser interface {
	abs() float64
}

func getAbs(value abser) float64 {
	return value.abs()
}

type floater float64

func (value floater) abs() float64 {
	if value < 0 {
		return -float64(value)
	}
	return float64(value)
}

type point struct {
	x float64
	y float64
}

func (value *point) abs() float64 {
	if value == nil {
		return 0
	}
	return math.Sqrt(value.x*value.x + value.y*value.y)
}

//implementing Stringer interface. A Stringer is a type that can describe itself as a string.
func (value *point) String() string {
	return fmt.Sprintf("x:%v|y:%v", value.x, value.y)
}

func describe(values []interface{}) {
	for _, value := range values {
		fmt.Printf("%v-%T, ", value, value)
	}
	fmt.Println()
}

func main() {

	a := floater(-5)
	b := point{2, 3}
	fmt.Println(getAbs(a), getAbs(&b))

	var c abser
	c = a
	fmt.Println(c)
	c = &b
	fmt.Println(c)

	var d point
	c = &d
	fmt.Println(d.abs(), c.abs())

	describe([]interface{}{1, 2.5, "text", false})

	assertions()

	e := point{5, 7}
	f := point{9, 5}
	fmt.Println(&e, &f)

}
