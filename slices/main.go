package main

import "fmt"

type point struct {
	x int
	y int
}

func describeSlice(name string, x []int) {
	fmt.Println(name, x, len(x), cap(x), x == nil)
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
	}{{1, false}, {2, true}}
	fmt.Println("slice5", slice5)

	slice6 := array[:0]
	describeSlice("slice6", slice6)

	//higher limit of a slice can be extended upto its "capacity".
	//changing the higher limit of a slice does not change its "capacity"
	slice6 = slice6[:4]
	describeSlice("slice6", slice6)

	//changing the lower limit of a slice "changes" its "capacity"
	slice6 = slice6[2:]
	describeSlice("slice6", slice6)

	var slice7 []int
	describeSlice("slice7", slice7)

	slice8 := make([]int, 5)
	describeSlice("slice8", slice8)

	slice9 := make([]int, 0, 5)
	describeSlice("slice9", slice9)
	slice9 = slice9[:cap(slice9)]
	describeSlice("slice9", slice9)
	slice9 = slice9[1:]
	describeSlice("slice9", slice9)
	slice9 = slice9[:0]
	describeSlice("slice9", slice9)
	slice9 = slice9[:2]
	describeSlice("slice9", slice9)
	slice9 = slice9[2:4]
	describeSlice("slice9", slice9)

	sliceOfSlicesOfStrings := [][]string{
		{"a", "b"},
		{"c", "d"},
	}
	fmt.Println(sliceOfSlicesOfStrings)

	var slice10 []int
	describeSlice("slice10", slice10)
	slice10 = append(slice10, 0, 2)
	describeSlice("slice10", slice10)
	slice10 = append(slice10, 4, 6, 8, 10, 12, 14, 16, 18, 20)
	describeSlice("slice10", slice10)

	for index, value := range slice10 {
		fmt.Printf("%d-%d, ", index, value)
	}
	fmt.Println()

	for _, value := range slice10 {
		fmt.Printf("%d, ", value)
	}
	fmt.Println()

	for index := range slice10 {
		fmt.Printf("%d, ", index)
	}
	fmt.Println()

}
