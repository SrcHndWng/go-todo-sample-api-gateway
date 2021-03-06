# About This

Todoを登録・削除・参照するAPIのサンプル。

Serverless Framework、API Gateway + Lambda、DynamoDbを使用。

## テストについて

* 現時点ではテストはDynamoDbを直に更新している。
* デプロイ時、「--stage」を指定ない場合は「dev」を指定した場合はテスト用テーブルを作成し、そこをテストでは更新する。
    
## Todoの採番について

* 採番についてはテーブル内のデータ数 + 1をしている。

# Compile & Deploy

0. Compile & Deploy

```
$ make build; serverless deploy;
```

1. Compile function

```
$ make build
```

2. Deploy!

```
serverless deploy
```

# Curl

```
$ curl -i -X POST https://api_gateway_domain/dev/todos -d 'Hello, world!'
$ curl -i https://api_gateway_domain/dev/todos
$ curl -i https://api_gateway_domain/dev/todos/2
$ curl -i -X PUT https://api_gateway_domain/dev/todos/4 -d 'this is update!'
$ curl -i -X DELETE https://api_gateway_domain/dev/todos/4
```
