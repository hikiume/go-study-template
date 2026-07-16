package main

import (
	"testing"
)

func TestGenerateToken(t *testing.T) {
	byteLength := 16
	expectedCharLength := byteLength * 2

	token1, err := GenerateToken(byteLength)
	if err != nil {
		t.Fatalf("GenerateToken でエラーが発生しました: %v", err)
	}

	if len(token1) != expectedCharLength {
		t.Errorf("GenerateToken() の文字数 = %d, 期待値 %d", len(token1), expectedCharLength)
	}

	token2, err := GenerateToken(byteLength)
	if err != nil {
		t.Fatalf("2回目の GenerateTokenでエラーが発生しました: %v", err)
	}

	if token1 == token2 {
		t.Errorf("2回続けて同じトークンが生成されました。乱数が機能していない可能性があります")
	}
}
