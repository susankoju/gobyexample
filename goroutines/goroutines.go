package main

import (
	"fmt"
	"time"
)

func printFrom(from string) {
	for i := range 4 {
		fmt.Println(from, ":", i)
	}
}

func main() {
	printFrom("1. Direct")

	go printFrom("2. Go Routines")

	printFrom("3. Direct 2")

	go printFrom("4. Go Routines 2")
	go printFrom("5. Go Routines 3")

	time.Sleep(time.Second)

	fmt.Println("Done")
}
