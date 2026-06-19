package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/posts/1"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("リクエストに失敗しました")
	}
	defer resp.Body.Close()

	fmt.Printf("Status Code: %d\n", resp.StatusCode)
}
