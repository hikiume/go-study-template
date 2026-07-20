package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthMiddleware(t *testing.T) {
	successHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	handlerToTest := authMiddleware(successHandler)

	tests := []struct {
		name           string
		requestHeader  string
		expectedStatus int
	}{
		{
			name:           "正常系: 正しいAPIキーならアクセス許可",
			requestHeader:  "secret123",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "異常系: 間違ったAPIキーなら認証エラー",
			requestHeader:  "wrong-key",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "異常系: APIキーが未指定なら認証エラー",
			requestHeader:  "",
			expectedStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/test", nil)
			if err != nil {
				t.Fatalf("リクエストの作成失敗: %v", err)
			}

			if tt.requestHeader != "" {
				req.Header.Set("X-API-KEY", tt.requestHeader)
			}

			rec := httptest.NewRecorder()

			handlerToTest.ServeHTTP(rec, req)

			if rec.Code != tt.expectedStatus {
				t.Errorf("ステータスコード = %d, 期待値 %d", rec.Code, tt.expectedStatus)
			}
		})
	}
}
