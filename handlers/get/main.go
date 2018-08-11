package main

import (
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/SrcHndWng/go-todo-sample-api-gateway/model/todo"
	"github.com/SrcHndWng/go-todo-sample-api-gateway/response"
)

// Handler is the only one entry point.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id, err := strconv.Atoi(request.PathParameters["id"])
	if err != nil {
		return response.Error(err)
	}
	data, err := todo.Get(id)
	if err != nil {
		return response.Error(err)
	}
	b, err := json.Marshal(data)
	if err != nil {
		return response.Error(err)
	}
	return response.Success(string(b))
}

func main() {
	lambda.Start(Handler)
}
