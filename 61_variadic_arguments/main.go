package main

import "fmt"

func Sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println(Sum())

	fmt.Println(Sum(10, 20))

	fmt.Println(Sum(1, 2, 3, 4, 5))

	slice := []int{5, 5, 5}
	fmt.Println(Sum(slice...))
}
