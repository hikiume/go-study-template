package main

import (
	"fmt"
)

// init()がmainより先に実行される
// 変数の多いコードで初期設定を行う時に利用する
func init() {
	fmt.Println("init !")
}

func main() {
	fmt.Printf("Hello world")
}
