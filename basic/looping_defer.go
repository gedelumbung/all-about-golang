package main

import "fmt"

func main() {
	fmt.Println("loading")

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
