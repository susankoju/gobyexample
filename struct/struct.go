package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{name: "susan", age: 1})
	fmt.Println(person{})

	s := person{"susan", 1}
	fmt.Println(s)
}
