package main

import (
	"fmt"
	"log"
	"os"
	"unicode/utf8"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/text"
)

func main() {

	file, err := os.ReadFile("example.md")
	if err != nil {
		log.Fatal("Not Found File")
	}

	source := file

	reader := text.NewReader(source)
	parser := goldmark.DefaultParser()
	doc := parser.Parse(reader)

	doc.Dump(source, 0)

	totalChars := utf8.RuneCount(source)
	fmt.Printf("全体の文字数%d 文字\n", totalChars)
}
