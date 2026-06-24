package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.ReadFile("example.json")

	if err != nil {
		log.Fatalln("ファイル読み込みエラー")
	}

	fmt.Println(string(file))

	var value map[string]string
	jsonErr := json.Unmarshal(file, &value)
	if jsonErr != nil {
		log.Fatalln("json Unmarshal")
	}

	fmt.Println(value["test"])
	fmt.Println(value["test_num"])
}
