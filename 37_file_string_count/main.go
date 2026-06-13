package main

import (
	"log"
	"os"
	"unicode/utf8"
)

func main() {
	// ファイルから文字を取得
	f, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal("ファイルが読み込みに失敗しました。")
	}

	// 文字のカウントをする
	count := utf8.RuneCountInString(string(f))

	// 表示する
	println(count)
}
