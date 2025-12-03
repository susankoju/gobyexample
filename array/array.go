package main

import "fmt"

func main() {
	var nums [3]int
	nums[0] = 1

	fmt.Println(nums)

	var nums2 [2]int = [2]int{1, 2}
	fmt.Printf("%v\n", nums2)

	var num3 = [...]int{1, 2, 3}
	fmt.Printf("%v\n", num3)

	var matrix = [2][3]int{{1, 2, 3}, {1, 2, 3}}

	fmt.Println(matrix)

	var names = [...]string{"Susan", "Koju"}
	fmt.Println(names)
}
