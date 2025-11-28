package main

import "fmt"

func main() {
	var a = 0
	for a < 10 {
		fmt.Println(a)
		a++
	}

	for {
		fmt.Println(a)
		a++
		if a > 20 {
			break
		}
	}

	for i := range 3 {
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
