package main

import "fmt"

func main() {
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("uninit:", s, s == nil, len(s) == 0, cap(s))

	s[0] = "Susan"

	fmt.Println(s)

	s = append(s, "1", "2", "3")

	fmt.Println(s)

}
