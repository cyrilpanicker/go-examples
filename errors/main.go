package main

import "fmt"

type divideByZeroError struct {
	message string
}

//implementing error interface
func (v divideByZeroError) Error() string {
	return fmt.Sprintf(v.message)
}

func divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, divideByZeroError{fmt.Sprintf("divide-by-zero-error : %v/%v", a, b)}
	}
	return a / b, nil
}

func main() {
	pairs := [][2]float64{
		{3, 2},
		{10, 5},
		{16, 4},
		{3, 0},
		{5, 2},
	}
	for _, pair := range pairs {
		result, error := divide(pair[0], pair[1])
		if error != nil {
			fmt.Println(error)
		} else {
			fmt.Println(result)
		}
	}
}
