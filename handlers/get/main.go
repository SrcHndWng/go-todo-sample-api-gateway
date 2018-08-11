package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/SrcHndWng/go-todo-sample-api-gateway/model"
	"github.com/SrcHndWng/go-todo-sample-api-gateway/response"
)

// Handler is the only one entry point.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var (
		id    int
		err   error
		exist bool
		data  model.Todo
		b     []byte
	)

	if id, err = strconv.Atoi(request.PathParameters["id"]); err != nil {
		return response.Error(err)
	}
	todo := model.NewTodo()
	if exist, err = todo.IsExist(id); err != nil {
		return response.Error(err)
	}
	if !exist {
		return response.NotFound(fmt.Sprintf("id = %v not found.", id))
	}
	if data, err = todo.Get(id); err != nil {
		return response.Error(err)
	}
	if b, err = json.Marshal(data); err != nil {
		return response.Error(err)
	}
	return response.Success(string(b))
}

func main() {
	lambda.Start(Handler)
}
