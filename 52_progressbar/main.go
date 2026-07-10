package main

import (
	"fmt"
	"time"

	"github.com/schollz/progressbar/v3"
)

func main() {
	maxTasks := 100

	bar := progressbar.Default(int64(maxTasks), "ダウンロード中")

	for range maxTasks {
		time.Sleep(20 * time.Millisecond)

		bar.Add(1)
	}

	fmt.Println("\nすべての処理が正常に完了しました")
}
