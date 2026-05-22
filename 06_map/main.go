package main

import (
	"fmt"
)

func main() {
	// sample1()
	// sample2()
	sample3()
}

func sample1() {
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	m["banana"] = 300
	fmt.Println(m)

	// マップ内に存在しないキーを指定して値を取り出そうとすると0が出力される
	fmt.Println(m["orange"])
}

// 変数1,変数2 :=マップ[キー] 要素が存在するかを確認できる
func sample2() {
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)

	v, ok := m["apple"]
	fmt.Println(v, ok)

	v2, ok2 := m["orange"]
	fmt.Println(v2, ok2)
}

// make()を使用して空のmapを作成してから値を入れることができる
func sample3() {
	m := make(map[string]int)
	m["pc"] = 5000
	fmt.Println(m)
}
