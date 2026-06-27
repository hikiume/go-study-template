package main

import "fmt"

func main() {
	original := []int{1, 2, 3}
	sub := original[0:2]

	sub[0] = 99

	// Goのスライスは元のデータの特定の位置を指し示しているだけ
	// そのため結果が[99 2 3]になる
	fmt.Println("original:", original)
	fmt.Println("sub:", sub)
}
