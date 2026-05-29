package main

import (
	"fmt"
	"os"
)

func main() {
	openFile()
}

func openFile() {
	file, _ := os.Open("./go.mod")
	defer file.Close()
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))
}

/*
最初に書いたdefer文の処理が最後に実行される
この場合だと
run
success
3
2
1
になる
*/
func sample() {
	fmt.Println("run")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("success")
}
