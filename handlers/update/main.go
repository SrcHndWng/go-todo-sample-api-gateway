package main

import (
	"strconv"

	"github.com/SrcHndWng/go-todo-sample-api-gateway/model/todo"
	"github.com/SrcHndWng/go-todo-sample-api-gateway/response"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler is the only one entry point.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id, err := strconv.Atoi(request.PathParameters["id"])
	if err != nil {
		return response.Error(err)
	}
	if err := todo.Update(id, request.Body); err != nil {
		return response.Error(err)
	}
	return response.Success("")
}

func main() {
	lambda.Start(Handler)
}
