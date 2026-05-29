package main

import (
	"fmt"

	"go-study/29_package/mylib"
	"go-study/29_package/mylib/under"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib.Average(s))

	mylib.Say()
	under.Hello()

	person :=mylib.Person{Name: "Mike",Age: 20}
	fmt.Println(person)

	fmt.Println(mylib.Public)
	// 小文字なのでエラーが発生する
	// fmt.Println(mylib.public)
}
