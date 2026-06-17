package main

import (
	"log"
	"os"

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
}
