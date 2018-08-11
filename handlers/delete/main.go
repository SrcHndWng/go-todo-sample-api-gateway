package main

import (
	"fmt"
	"strconv"

	"github.com/SrcHndWng/go-todo-sample-api-gateway/model"
	"github.com/SrcHndWng/go-todo-sample-api-gateway/response"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler is the only one entry point.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var (
		id    int
		err   error
		exist bool
	)

	if id, err = strconv.Atoi(request.PathParameters["id"]); err != nil {
		return response.Error(err)
	}
	todo := model.Todo{}
	if exist, err = todo.IsExist(id); err != nil {
		return response.Error(err)
	}
	if !exist {
		return response.NotFound(fmt.Sprintf("id = %v not found.", id))
	}
	if err = todo.Delete(id); err != nil {
		return response.Error(err)
	}
	return response.Success("")
}

func main() {
	lambda.Start(Handler)
}
