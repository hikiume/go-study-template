package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	numbers := []int{5, 2, 6, 3, 1, 4}
	slices.Sort(numbers)
	fmt.Println("数値のソート後:", numbers)

	fruites := []string{"banana", "apple", "orrage", "grape"}
	slices.Sort(fruites)
	fmt.Println("文字列のソート後:", fruites)

	SortDesc(numbers)
}

func SortDesc(nums []int) {
	slices.SortFunc(nums, func(a, b int) int {
		return cmp.Compare(b, a)
	})
}
