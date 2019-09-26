package main

import "fmt"

var boolean1, boolean2, boolean3 bool
var int1, int2 int = 1, 2
var int3, int4 = 3, 4
var (
	float1   = 3.5
	boolean4 bool
	string1  string = "string1"
)

const (
	pi = 3.142
)
const truth = true

func main() {
	var boolean5 bool
	var string2, boolean6 = "string2", true
	string3, boolean7 := "string3", false
	var (
		string4  = "string4"
		boolean8 bool
		string5  string = "string5"
	)
	const constant1 = "constant1"
	const (
		constant2 = "constant2"
		constant3 = false
	)
	fmt.Println(boolean1, boolean2, boolean3, boolean4, boolean5,
		boolean6, boolean7, boolean8, string1, string2, string3, string4, string5)
	fmt.Println(pi, truth, constant1, constant2, constant3)
}
