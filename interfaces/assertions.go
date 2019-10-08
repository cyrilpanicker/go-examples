package main

import "fmt"

func assert(values []interface{}) {
	for _, value := range values {
		if v, ok := value.(string); ok {
			fmt.Printf("string-%v, ", v)
		} else if v, ok := value.(float64); ok {
			fmt.Printf("float-%v, ", v)
		} else {
			fmt.Printf("unknown-%v, ", value)
		}
	}
	fmt.Println()
}

func assertWithSwitch(values []interface{}) {
	for _, value := range values {
		switch value.(type) {
		case string:
			fmt.Printf("string-%v, ", value)
		case float64:
			fmt.Printf("float-%v, ", value)
		default:
			fmt.Printf("unknown-%v, ", value)
		}
	}
	fmt.Println()
}

func assertions() {
	fmt.Println("-----assertions start------")
	assert([]interface{}{"text", 3.14, "text1", 4.7, false})
	assertWithSwitch([]interface{}{"text", 3.14, "text1", 4.7, false})
	fmt.Println("-----assertions end------")
}
