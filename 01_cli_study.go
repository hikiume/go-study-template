package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("【エラー】名前を入力してください。例: go run 01_cli_study.go タロウ")
		return
	}

	name := args[1]

	fmt.Println("%s さん CLIは動作しています\n", name)
}
