FROM golang:1.21.5-alpine3.18
# alpineベースのイメージにはgitやbashが含まれていないため、インストールの必要がある
RUN apk update && apk add git

ENV TZ /usr/share/zoneinfo/Asia/Tokyo

WORKDIR /app

COPY /app/* ./

# go.modにあるパッケージの事前ダウンロード
RUN go mod download

# ビルド済みの実行可能なバイナリをインストール
RUN go install github.com/cosmtrek/air@latest

EXPOSE 5050

CMD ["air", "-c", ".air.toml"]