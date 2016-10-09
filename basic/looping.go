package main

import (
	"fmt"
)

func main() {
	fmt.Println("Standar looping ")
	for i := 0; i < 6; i++ {
		fmt.Println(i)
	}

	fmt.Println("Looping with break")
outerLoop:
	for i := 0; i < 7; i++ {
		for j := 0; j < 6; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Println("[", i, "][", j, "]")
		}
	}
}
