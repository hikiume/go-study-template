package main

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "my_secure_password"

	hash, err := HashPasswrod(password)

	if err != nil {
		t.Fatalf("HashPasswrod でエラーが発生しました: %v", err)
	}

	if hash == "" {
		t.Error("ハッシュ値が空文字です")
	}

	hash2, err := HashPasswrod(password)
	if err != nil {
		t.Fatalf("2回目の HashPasswrod でエラーが発生しました: %v", err)
	}

	if hash == hash2 {
		t.Error("同じパスワードから同じハッシュが生成されました。ソルトが機能していません。")
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password := "correct_password_123"
	hash, err := HashPasswrod(password)
	if err != nil {
		t.Fatalf("テスト用のハッシュ作成に失敗しました: %v", err)
	}

	tests := []struct {
		name          string
		inputPassword string
		want          bool
	}{
		{
			name:          "正しいパスワードでの認証成功",
			inputPassword: "correct_password_123",
			want:          true,
		},
		{
			name:          "間違ったパスワードで認証失敗",
			inputPassword: "wrong_password",
			want:          false,
		},
		{
			name:          "空文字のパスワードで認証失敗",
			inputPassword: "",
			want:          false,
		},
		{
			name:          "大文字小文字が異なるパスワードで認証失敗",
			inputPassword: "Correct_Password_123",
			want:          false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckPasswordHash(tt.inputPassword, hash)
			if got != tt.want {
				t.Errorf("CheckPasswordHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
