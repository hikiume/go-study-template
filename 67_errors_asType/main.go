package main

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

func validate(name string) error {
	if name == "" {
		return &ValidationError{
			Field: "name", Message: "必須項目です",
		}
	}
	return nil
}

func main() {
	err := validate("")

	// errors.AsTypeはエラーチェーンの中から特定の型のエラーを取り出す
	if ve, ok := errors.AsType[*ValidationError](err); ok {
		fmt.Printf("バリデーションエラー: %s\n", ve.Field)
	}
}
