package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println("処理を開始します")

	time.Sleep(2 * time.Second)

	fmt.Println("処理が完了しました。")

	duration := time.Since(start)

	fmt.Printf("Done in %.2f scondes\n", duration.Seconds())
}
