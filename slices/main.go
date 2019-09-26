package main

import "fmt"

type point struct {
	x int
	y int
}

func main() {

	array := [6]int{0, 1, 2, 3, 4, 5}
	fmt.Println("array", array)

	var slice1 []int = array[1:4]
	slice2 := array[2:5]
	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2)

	slice1[2] = 99
	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2)
	fmt.Println("array", array)

	slice3 := []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("slice3", slice3)
	fmt.Println("slice3-slice", slice3[1:])
	fmt.Println("slice3-slice", slice3[:5])

	slice4 := []point{{1, 2}, {3, 4}}
	fmt.Println("slice4", slice4)

	slice5 := []struct {
		a int
		b bool
	}{
		{1, false},
		{2, true},
	}
	fmt.Println("slice5", slice5)

}
