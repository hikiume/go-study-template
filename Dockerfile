# 軽量なAlpine LinuxベースのGoイメージを使用
FROM golang:1.24-alpine

# SQLiteのビルドに必要なツール（GCC、MUSL開発パック）をインストール
RUN apk add --no-cache gcc musl-dev sqlite

# コンテナ内の作業ディレクトリを設定
WORKDIR /app

# CGOを有効化する
ENV CGO_ENABLED=1

# コンテナが起動した時のデフォルトのコマンド（シェルを起動）
CMD ["sh"]
