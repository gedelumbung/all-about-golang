package main

import "fmt"

func fact(n int) int {
	if n < 3 {
		return 2
	}
	return n * fact(n-2)
}

func main() {
	fmt.Println(fact(10))
}
