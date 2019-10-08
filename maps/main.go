package main

import "fmt"

type point struct {
	x int
	y int
}

func main() {

	map1 := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
		"key4": "value4",
		"key5": "value5",
	}
	fmt.Println("map1", map1)

	map2 := map[string]struct {
		x int
		y int
	}{
		"key1": {x: 1, y: 2},
		"key2": {x: 3, y: 4},
		"key3": {x: 5, y: 6},
		"key4": {x: 7, y: 8},
	}
	fmt.Println("map2", map2)

	map3 := map[string]point{
		"point1": {x: 1, y: 2},
		"point2": {x: 3, y: 4},
		"point3": {x: 5, y: 6},
		"point4": {x: 7, y: 8},
	}
	fmt.Println("map3", map3)

	map4 := map[string]string{}
	fmt.Println("map4", map4, map4 == nil, len(map4))
	map4["key1"] = "value1"
	fmt.Println("map4", map4)

	map5 := make(map[string]string)
	fmt.Println("map5", map5, map5 == nil, len(map5))
	map5["key1"] = "value1"
	fmt.Println("map5", map5)

	map6 := make(map[string]point)
	map6["key1"] = point{1, 2}
	fmt.Println("map6", map6)

	fmt.Println(map5["key1"])
	map5["key1"] = "value2"
	fmt.Println(map5["key1"])
	value, isPresent := map5["key1"]
	fmt.Println(value, isPresent)
	delete(map5, "key1")
	value1, isPresent1 := map5["key1"]
	fmt.Println(value1, isPresent1)

	map7 := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	map8 := map7
	//map/slice/channel values are pointers wrapped by struct values
	//the destination and source map/slice/channel values in an assignment will share the same underlying data pointed by the wrapped pointer
	map8["key2"] = "value99"
	fmt.Println(map7, map8)

}
