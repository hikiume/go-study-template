package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "go python go rust go python java"
	words := strings.Fields(text)

	counts := make(map[string]int)

	for _, word := range words {
		counts[word]++
	}

	fmt.Println("単語の集計")
	for word, count := range counts {
		fmt.Printf("%s: %d回\n", word, count)
	}
}
