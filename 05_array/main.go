package main

import (
	"fmt"
)

func main() {
	// makeSample()
	byteSample()
}

func makeSample() {

	// makeで配列を定義 型,長さ,容量を定義できる
	n := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)

	n = append(n, 0, 0)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)

	n = append(n, 1, 2, 3, 4, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)

	// 省略すると長さと容量が同じになる
	a := make([]int, 3)
	fmt.Printf("len=%d cap=%d value=%v\n", len(a), cap(a), a)
}

func byteSample() {
	b := []byte{72, 73}
	fmt.Println(b)
	fmt.Println(string(b))

	c := []byte("HI")
	fmt.Println(c)
}
