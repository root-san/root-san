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

生成されたコードのインターフェースを満たすようにapp/handlerでServerのメソッドを実装する

## example of api

### POST /rooms

```
curl -XPOST -H "Content-Type:application/json" -d '{"id": "c3cf4b9a-7316-4a41-bf60-194080cbe682", "name": "hoghoge"}' localhost:8080/rooms
```

### POST /rooms/{roomId}/members

```
curl -XPOST \
-H "Content-Type:application/json" \
-d '{"id": "c3cf4b9a-7316-4a41-bf60-194080cbe682", "name": "hoghoge"}' localhost:8080/rooms/c3cf4b9a-7316-4a41-bf60-194080cbe682/members
```

### DELETE /rooms/{roomId}/members/{memberId}

```
curl -XDELETE \
-H "Content-Type:application/json" \
localhost:8080/rooms/c3cf4b9a-7316-4a41-bf60-194080cbe682/members/c3cf4b9a-7316-4a41-bf60-194080cbe682
```

### POST /rooms/{roomId}/events
eventAt
```
curl -XPOST \
-H 'Content-Type:application/json' \
-d '{
  "id": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
  "name": "string",
  "eventType": "outer",
  "amount": 0,
  "eventAt": "2021-05-31T16:27:35+09:00",
  "txns": [
    {
      "amount": 0,
      "receiver": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
      "payer": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
      "id": "3fa85f64-5717-4562-b3fc-2c963f66afa6"
    }
  ]
}' \
localhost:8080/rooms/c3cf4b9a-7316-4a41-bf60-194080cbe682/events
```

### DELETE /rooms/{roomId}/events/{eventId}

```
curl -XDELETE \
-H "Content-Type:application/json" \
localhost:8080/rooms/c3cf4b9a-7316-4a41-bf60-194080cbe682/events/3fa85f64-5717-4562-b3fc-2c963f66afa6
```

### PUT /rooms/{roomId}/events/{eventId}

```
curl -XPUT \
-H 'Content-Type:application/json' \
-d '{
  "id": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
  "name": "string",
  "eventType": "outer",
  "amount": 0,
  "eventAt": "2021-05-31T16:27:35+09:00",
  "txns": [
    {
      "amount": 0,
      "receiver": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
      "payer": "3fa85f64-5717-4562-b3fc-2c963f66afa7",
      "id": "3fa85f64-5717-4562-b3fc-2c963f66afa6"
    }
  ]
}' \
localhost:8080/rooms/c3cf4b9a-7316-4a41-bf60-194080cbe682/events/3fa85f64-5717-4562-b3fc-2c963f66afa6
```

