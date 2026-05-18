# 軽量なAlpine LinuxベースのGoイメージを使用
FROM golang:1.24-alpine

# コンテナ内の作業ディレクトリを設定
WORKDIR /app

# コンテナが起動した時のデフォルトのコマンド（シェルを起動）
CMD ["sh"]
