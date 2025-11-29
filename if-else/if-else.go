package main

import "fmt"

func main() {
	num := 5
	if num%3 == 0 {
		fmt.Println("Divisible by 3")
	} else {
		fmt.Println("Not divisible by 3")
	}

	if num%5 == 0 || num%3 == 0 {
		fmt.Println("Is divisible by 5 or 3")
	}

	if i := 0; i == 0 {
		fmt.Println("i is 0")
	} else if i > 10 {
		fmt.Println("i is greater than 10")
	}
}
