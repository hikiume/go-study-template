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
