package main

import (
	"errors"
	"fmt"
	"strings"
)

type User struct {
	Email    string
	Password string
}

func ValidateUser(user User) error {
	if !strings.Contains(user.Email, "@") || strings.HasPrefix(user.Email, "@") || strings.HasSuffix(user.Email, "@") {
		return errors.New("無効なメールアドレスの形式です")
	}

	if len(user.Password) < 8 {
		return errors.New("パスワードは8文字以上で入力してください")
	}

	return nil
}

func main() {
	goodUser := User{Email: "test@example.com", Password: "super_secret_123"}
	if err := ValidateUser(goodUser); err != nil {
		fmt.Printf("登録失敗: %v\n", err)
	} else {
		fmt.Printf("ユーザーA: バリデーション成功!\n")
	}

	badUser := User{Email: "invalid-email", Password: "short"}
	if err := ValidateUser(badUser); err != nil {
		fmt.Printf("ユーザーB: 登録失敗 -> %v\n", err)
	}
}
