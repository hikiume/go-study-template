package main

import (
	"fmt"
	"net/http"
	"os"
)

type healthHandler struct{}

func (h healthHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("GET /health", healthHandler{})

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Fprintf(os.Stderr, "サーバーエラー: %v\n", err)
		os.Exit(1)
	}
}
