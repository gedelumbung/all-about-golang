package main

import (
	"fmt"
)

func main() {
	var point byte = 6

	if point > 7 {
		fmt.Printf("Anda lulus, nilai anda %d \n", point)
	} else {
		fmt.Printf("Anda tidak lulus, nilai anda %d \n", point)
	}
}
