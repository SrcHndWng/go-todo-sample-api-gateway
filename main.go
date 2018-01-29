package main

import (
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"./api/todos"
)

// Handler is the only one entry point.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch request.Path {
	case "/todos":
		switch request.HTTPMethod {
		case "POST":
			err := todos.Create(request.Body)
			if err != nil {
				return events.APIGatewayProxyResponse{Body: "", StatusCode: 500}, err
			}
			return events.APIGatewayProxyResponse{Body: "post!", StatusCode: 200}, nil
		case "GET":
			todos, err := todos.List()
			if err != nil {
				return events.APIGatewayProxyResponse{Body: "", StatusCode: 500}, err
			}
			return events.APIGatewayProxyResponse{Body: fmt.Sprintf("%v", todos), StatusCode: 200}, nil
		}
	}
	return events.APIGatewayProxyResponse{Body: "", StatusCode: 500}, errors.New("Unhandled Request")
}

func main() {
	lambda.Start(Handler)
}
