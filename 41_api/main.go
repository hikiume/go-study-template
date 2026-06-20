package main

import (
	"fmt"
	"io"
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ボディの読み込みに失敗しました")
	}

	fmt.Printf("Respons Body:%s\n", string(body))

}
