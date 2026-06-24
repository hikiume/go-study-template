package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

type Bookstore struct {
	XMLName xml.Name `xml:"bookstore"`
	Books   []Book   `xml:"book"`
}

type Book struct {
	ID     string  `xml:"id,attr"`
	Title  string  `xml:"title"`
	Author string  `xml:"author"`
	Price  float64 `xml:"price"`
}

func main() {
	xmlData, readErr := os.ReadFile("example.xml")
	if readErr != nil {
		log.Fatalln("not file open")
	}

	var store Bookstore
	err := xml.Unmarshal(xmlData, &store)
	if err != nil {
		log.Fatalln("XMLのパースに失敗")
	}

	fmt.Printf("本屋のデータ: %+v\n\n", store)
	for _, book := range store.Books {
		fmt.Printf("ID : %s | タイトル: %s | 著者: %s | 価格: %.2f\n", book.ID, book.Title, book.Author, book.Price)
	}
}
