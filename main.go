package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"./api/todos"
)

// Handler is the only one entry point.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch request.Resource {
	case "/todos":
		switch request.HTTPMethod {
		case "POST":
			err := todos.Create(request.Body)
			if err != nil {
				return errorResponse(err)
			}
			return successResponse("")
		case "GET":
			todos, err := todos.List()
			if err != nil {
				return errorResponse(err)
			}
			data, err := json.Marshal(todos)
			if err != nil {
				return errorResponse(err)
			}
			return successResponse(string(data))
		}
	case "/todos/{id}":
		// fmt.Printf("PathParameters[:id] = %v\n", request.PathParameters["id"])
		// return successResponse("/todos/{id}")
		id, err := strconv.Atoi(request.PathParameters["id"])
		if err != nil {
			return errorResponse(err)
		}
		todo, err := todos.Get(id)
		if err != nil {
			return errorResponse(err)
		}
		data, err := json.Marshal(todo)
		if err != nil {
			return errorResponse(err)
		}
		return successResponse(string(data))
	}
	return errorResponse(errors.New("Unhandled Request"))
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
