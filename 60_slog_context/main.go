package main

import (
	"context"
	"log/slog"
	"os"
)

type contextKey string

const requestIDKey contextKey = "request_id"

func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey, requestID)
}

func LogWithContext(ctx context.Context, msg string, attrs ...slog.Attr) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	reqID, ok := ctx.Value(requestIDKey).(string)
	if !ok {
		reqID = "unknown"
	}

	args := append([]slog.Attr{slog.String("request_id", reqID)}, attrs...)

	logger.LogAttrs(ctx, slog.LevelInfo, msg, args...)
}

func ProcessOrder(ctx context.Context, itemID string, amount int) {
	LogWithContext(ctx, "注文を開始しました",
		slog.String("item_id", itemID),
		slog.Int("amount", amount))
}

func main() {
	ctxA := WithRequestID(context.Background(), "req_12345")
	ProcessOrder(ctxA, "item_999", 2)

	ctxB := WithRequestID(context.Background(), "req_67890")
	ProcessOrder(ctxB, "item_111", 1)
}
