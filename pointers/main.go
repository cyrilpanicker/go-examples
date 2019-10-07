package main

import "fmt"

func main() {

	i := 1
	p := &i
	fmt.Println(*p)
	*p = 3
	fmt.Println(*p)

}
