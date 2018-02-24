package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/SrcHndWng/go-todo-sample-api-gateway/api/todos"
	"github.com/SrcHndWng/go-todo-sample-api-gateway/model/todo"
)

// Handler is the only one entry point.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var id int
	var datas []todo.Todo
	var data todo.Todo
	var b []byte
	var err error

	switch request.Resource {
	case "/todos":
		switch request.HTTPMethod {
		case "POST":
			if err = todos.Create(request.Body); err != nil {
				return errorResponse(err)
			}
			return successResponse("")
		case "GET":
			if datas, err = todos.List(); err != nil {
				return errorResponse(err)
			}
			if b, err = json.Marshal(datas); err != nil {
				return errorResponse(err)
			}
			return successResponse(string(b))
		}
	case "/todos/{id}":
		switch request.HTTPMethod {
		case "GET":
			if id, err = strconv.Atoi(request.PathParameters["id"]); err != nil {
				return errorResponse(err)
			}
			if data, err = todos.Get(id); err != nil {
				return errorResponse(err)
			}
			if b, err = json.Marshal(data); err != nil {
				return errorResponse(err)
			}
			return successResponse(string(b))
		case "PUT":
			if id, err = strconv.Atoi(request.PathParameters["id"]); err != nil {
				return errorResponse(err)
			}
			if err = todos.Update(id, request.Body); err != nil {
				return errorResponse(err)
			}
			return successResponse("")
		case "DELETE":
			if id, err = strconv.Atoi(request.PathParameters["id"]); err != nil {
				return errorResponse(err)
			}
			if err = todos.Delete(id); err != nil {
				return errorResponse(err)
			}
			return successResponse("")
		}
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
