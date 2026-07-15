package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func GenerateToken(length int) (string, error) {
	b := make([]byte, length)

	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	// バイト配列を16進数の文字列に変換して返す (例: [255, 0] -> "ff00")
	return hex.EncodeToString(b), nil
}

func main() {
	token, err := GenerateToken(16)
	if err != nil {
		fmt.Printf("トークン生成失敗: %v\n", err)
		return
	}

	fmt.Printf("生成された安全なトークン: %s\n", token)
	fmt.Printf("トークンの文字数: %d文字\n", len(token))
}
