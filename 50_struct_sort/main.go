package main

import (
	"cmp"
	"fmt"
	"slices"
)

type User struct {
	Name string
	Age  int
}

func main() {
	sort()

	multiSort()
}

func sort() {
	users := []User{
		{Name: "Taro", Age: 25},
		{Name: "Hanako", Age: 20},
		{Name: "Jiro", Age: 30},
	}

	slices.SortFunc(users, func(a, b User) int {
		return cmp.Compare(a.Age, b.Age)
	})

	fmt.Println("年齢が低い順:", users)
}

// 複数の条件でソート
func multiSort() {
	users := []User{
		{Name: "Taro", Age: 25},
		{Name: "Hanako", Age: 20},
		{Name: "Alice", Age: 25},
	}

	slices.SortFunc(users, func(a, b User) int {
		result := cmp.Compare(a.Age, b.Age)

		if result != 0 {
			return result
		}

		return cmp.Compare(a.Name, b.Name)
	})

	fmt.Println(users)
}
