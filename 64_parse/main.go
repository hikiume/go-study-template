package main

import (
	"fmt"
	"time"
)

func main() {
	dateStr := "2026-07-29"

	layout := "2006-01-02"

	t, err := time.Parse(layout, dateStr)
	if err != nil {
		fmt.Println("パースエラー:", err)
		return
	}

	fmt.Println("パース結果(UTC):", t)
	fmt.Println("年:", t.Year())
	fmt.Println("月:", t.Month())
	fmt.Println("日:", t.Day())
}
