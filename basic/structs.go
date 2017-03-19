package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	fmt.Println(person{"Wayan", 10})

	fmt.Println(person{name: "Made", age: 15})

	fmt.Println(person{name: "Nyoman"})

	fmt.Println(&person{name: "Ketut", age: 30})

	s := person{name: "Putu", age: 42}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 25
	fmt.Println(sp.age)
}
