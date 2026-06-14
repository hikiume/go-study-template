package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.ReadDir("go-study")
	if err != nil {
		log.Fatal("フォルダの読み込みに失敗")
	}

	for _, v := range file {
		println(v.Name())
	}
}
