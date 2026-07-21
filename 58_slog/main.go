package main

import (
	"log/slog"
	"os"
)

func main() {
	// 出力フォーマットを「JSON」に設定する
	logger:=slog.New(slog.NewJSONHandler(os.Stdout,nil))

	// 普通のログ出力 (キーと値のペアで書く)
	logger.Info("ユーザーがログインしました","user_id",101,"ip_address","192.168.1.1",)

	// エラーログの出力
	logger.Error("データベース接続に失敗しました","db_name","users_db","retry_count",3)
}
