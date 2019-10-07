package main

import "fmt"

type point struct {
	x int
	y int
}

func main() {

	var array1 [2]int
	fmt.Println(array1, array1[0], array1[1])
	array1[1] = 1
	fmt.Println(array1, array1[0], array1[1])

	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	array3 := [2]point{{1, 2}, {3, 4}}
	fmt.Println(array3)

	array4 := [2]struct {
		a int
		b bool
	}{
		{1, true},
		{2, false},
	}
	fmt.Println(array4)

	array5 := [2]struct {
		x int
		y int
	}{
		{x: 1, y: 2},
		{x: 3, y: 4},
	}
	//creates a new copy of the array
	array6 := array5
	fmt.Println(array5, array6)
	array6[1].x = 99
	fmt.Println(array5, array6)

}
