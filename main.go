package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"./api/todos"
)

// Handler is the only one entry point.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	str := fmt.Sprintf("method:%s\npath:%s\nbody:%s\n", request.HTTPMethod, request.Path, request.Body)
	fmt.Println(str)

	var err error

	switch request.Path {
	case "/todos":
		switch request.HTTPMethod {
		case "POST":
			err = todos.Create(request.Body)
		case "GET":
			todos.List()
		}
	}

	if err != nil {
		return events.APIGatewayProxyResponse{Body: "", StatusCode: 500}, err
	}

	return events.APIGatewayProxyResponse{Body: str, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
