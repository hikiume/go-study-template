package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Date(2026, time.July, 30, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2026, time.December, 25, 0, 0, 0, 0, time.UTC)

	if t1.After(t2) {
		fmt.Println("t1はt2より未来の日時です")
	}

	if t1.Before(t2) {
		fmt.Println("t1はt2より過去の日時です")
	}

	if t1.Equal(t2) {
		fmt.Println("t1とt2は同じ日時です")
	}
}
