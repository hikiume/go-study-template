package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `josn:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/posts/1"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("リクエストに失敗しました")
	}
	defer resp.Body.Close()

	fmt.Printf("Status Code: %d\n", resp.StatusCode)

	var post Post

	if err := json.NewDecoder(resp.Body).Decode(&post); err != nil {
		log.Fatalf("デコード失敗: %v", err)
	}

	fmt.Printf("ID: %d\n", post.ID)
	fmt.Printf("タイトル: %s\n", post.Title)
}
