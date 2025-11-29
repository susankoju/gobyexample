package main

import (
	"fmt"
	"time"
)

func main() {
	num := 3

	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
		fmt.Println("Below 3")
	default:
		fmt.Println("Other")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	returnSame := func(a interface{}) interface{} {
		switch {
		case a == 3:
			return 3
		}
		return 0
	}

	fmt.Println(returnSame(1))
}
