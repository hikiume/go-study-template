package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	stdin := os.Stdin
	scanner := bufio.NewScanner(stdin)
	scanner.Scan()

	file, err := os.ReadDir(scanner.Text())
	if err != nil {
		log.Fatal("フォルダの読み込みに失敗")
	}

	for _, v := range file {
		println(v.Name())
	}
}
