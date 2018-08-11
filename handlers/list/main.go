package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/SrcHndWng/go-todo-sample-api-gateway/model"
	"github.com/SrcHndWng/go-todo-sample-api-gateway/response"
)

// Handler is the only one entry point.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var (
		datas []model.Todo
		err   error
		b     []byte
	)

	todo := model.Todo{}
	if datas, err = todo.List(); err != nil {
		return response.Error(err)
	}
	if b, err = json.Marshal(datas); err != nil {
		return response.Error(err)
	}
	return response.Success(string(b))
}

func main() {
	lambda.Start(Handler)
}
