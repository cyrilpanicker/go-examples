package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("Hello, Reader!")
	bytes := make([]byte, 4)
	for {
		n, err := reader.Read(bytes)
		if err == io.EOF {
			break
		}
		fmt.Printf("%q\n", bytes[:n])
	}
}
