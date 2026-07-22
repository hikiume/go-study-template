package main

import (
	"context"
	"fmt"
)

type KeyType string

const userKey KeyType = "current_user"

func main() {
	ctx := context.Background()

	ctwWithUser := context.WithValue(ctx, userKey, "Alice")

	Controller(ctwWithUser)
}

func Controller(ctx context.Context) {
	Service(ctx)
}

func Service(ctx context.Context) {
	if userName, ok := ctx.Value(userKey).(string); ok {
		fmt.Printf("[Service層] Contextから受け取ったユーザー名: %s\n", userName)
	} else {
		fmt.Println("[Service層] ユーザー情報が見つかりませんでした")
	}
}
