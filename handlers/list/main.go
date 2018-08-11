package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/SrcHndWng/go-todo-sample-api-gateway/model/todo"
)

// Handler is the only one entry point.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var datas []todo.Todo
	var b []byte
	var err error

	if datas, err = todo.List(); err != nil {
		return errorResponse(err)
	}
	if b, err = json.Marshal(datas); err != nil {
		return errorResponse(err)
	}
	return successResponse(string(b))
}

func successResponse(body string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: body}, nil
}

func errorResponse(err error) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("%+v\n", err)
	return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Internal Server Error!"}, nil
}

func main() {
	lambda.Start(Handler)
}
