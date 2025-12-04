package main

import "fmt"

type rect struct {
	width  int
	height int
}

func (r *rect) area() int {
	return r.height * r.width
}

func (r rect) perim() int {
	return 2*r.height + 2*r.width
}

func main() {
	r := rect{2, 3}
	fmt.Println("Area:", r.area())
	fmt.Println("Perimeter:", r.perim())
}
