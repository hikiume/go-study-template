package mylib

import "fmt"

// Public 大文字だと他のパッケージから呼び出せる
var Public string = "Public"

// 小文字だと他のパッケージから呼び出せない
var private string = "private"

type Person struct {
	Name string
	Age  int
}

func Say() {
	fmt.Println("Human!")
}
