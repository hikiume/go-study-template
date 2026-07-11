package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPasswrod(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	rawPassword := "super_secret_password_123"
	fmt.Printf("元のパスワード: %s\n", rawPassword)

	hashedPassword, err := HashPasswrod(rawPassword)
	if err != nil {
		log.Fatalf("ハッシュ化に失敗しました: %v", err)
	}

	fmt.Printf("DBに保存するハッシュ値: %s\n", hashedPassword)

	fmt.Println("\n--- ログイン時の検証シュミレーション ---")

	success := CheckPasswordHash("super_secret_password_123", hashedPassword)
	fmt.Printf("正しいパスワードでの認証結果: %t\n", success)

	fail := CheckPasswordHash("wrong_password", hashedPassword)
	fmt.Printf("間違ったパスワードでの認証結果: %t\n", fail)
}
