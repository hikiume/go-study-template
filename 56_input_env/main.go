package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Port   int
	APIURL string
}

func LoadConfig() (Config, error) {
	portStr := os.Getenv("APP_PORT")
	apiURL := os.Getenv("API_URL")

	port := 8080
	if portStr != "" {
		var err error
		port, err = strconv.Atoi(portStr)
		if err != nil {
			return Config{}, errors.New("APP_PORT は数値で指定してください")
		}
	}

	if apiURL == "" {
		return Config{}, errors.New("API_URL は必須項目です")
	}

	return Config{
		Port:   port,
		APIURL: apiURL,
	}, nil
}

func main() {
	fmt.Println("=== 正常系のシュミレーション ===")

	os.Setenv("APP_PORT", "3000")
	os.Setenv("API_URL", "https://api.example.com")

	cfg, err := LoadConfig()

	if err != nil {
		fmt.Printf("設定の読み込み失敗: %v\n", err)
	} else {
		fmt.Printf("読み込み成功! 起動ポート: %d, 接続先: %s\n", cfg.Port, cfg.APIURL)
	}

	fmt.Println("\n=== 異常系（エラー）のシュミレーション ===")
	os.Setenv("APP_PORT", "not-a-number")
	_, err = LoadConfig()
	if err != nil {
		fmt.Printf("期待値通りのエラー: %v\n", err)
	}
}
