package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Go Backend World!")
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Printf("[前処理] リクエスト受信: %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		log.Printf("[後処理] 処理完了。所要時間: %s", time.Since(start))
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-KEY")

		if apiKey != "secret123" {
			http.Error(w, "Unauthorized: 無効なAPIキーです", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(helloHandler)

	wrappedHandler := loggingMiddleware(authMiddleware(finalHandler))

	mux.Handle("/api/hello", wrappedHandler)

	fmt.Println("サーバーをポート 8080 で起動中... (Ctrl+C で停止)")

	log.Fatal(http.ListenAndServe(":8080", mux))
}
