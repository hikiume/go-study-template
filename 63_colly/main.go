package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	// コレクターの初期化
	c := colly.NewCollector(colly.AllowedDomains("example.com"))

	// HTML要素 (h1タグなど)を見つけた時の処理
	c.OnHTML("h1", func(e *colly.HTMLElement) {
		fmt.Println("H1テキスト:", e.Text)
	})

	// リクエスト実行時のログ表示
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// スクレイピング実行
	c.Visit("https://example.com")
}
