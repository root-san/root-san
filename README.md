# root-san

## Prerequisites

- Go: v1.18
- MySQL: 8.x

## set up

```sh
make dev
```

## 開発の流れ

oapi-codegenを利用し、APIスキーマ駆動で開発を行う

OASを編集し、エンドポイントを追加し以下のコマンドを実行すると、gen/api以下にAPIサーバーコード(server.gen.go)が生成される

```sh
make gen-api
```

生成されたコードのインターフェースを満たすようにapp/handlerのhandlerを実装する
