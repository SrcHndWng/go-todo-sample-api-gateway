# About This

Todoを登録・削除・参照するAPIのサンプル。

Serverless Framework、API Gateway + Lambda、DynamoDbを使用。

# Compile & Deploy

0. Compile & Deploy

```
$ GOOS=linux go build -o bin/main; serverless deploy;
```

1. Compile function

```
$ GOOS=linux go build -o bin/main
```
2. Deploy!

```
serverless deploy
```
