package main

import (
	"github.com/SrcHndWng/go-todo-sample-api-gateway/model"
	"github.com/SrcHndWng/go-todo-sample-api-gateway/response"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler is the only one entry point.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	todo := model.Todo{}
	if err := todo.Create(request.Body); err != nil {
		return response.Error(err)
	}
	return response.Success("")
}

func main() {
	lambda.Start(Handler)
}
