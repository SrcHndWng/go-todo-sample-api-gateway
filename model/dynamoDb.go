package model

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

// DynamoDb creates DynamoDb Object.
func DynamoDb() *dynamo.DB {
	return dynamo.New(session.New(), &aws.Config{Region: aws.String("ap-northeast-1")})
}

// Table create Table Object.
func Table(table string) dynamo.Table {
	db := DynamoDb()
	return db.Table(table)
}
