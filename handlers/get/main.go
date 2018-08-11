package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/SrcHndWng/go-todo-sample-api-gateway/model/todo"
	"github.com/SrcHndWng/go-todo-sample-api-gateway/response"
)

/*
type handler struct {
	err error
}

func (h handler) Atoi(request events.APIGatewayProxyRequest) (id int) {
	if h.err != nil {
		return
	}
	id, h.err = strconv.Atoi(request.PathParameters["id"])
	return
}

func (h handler) IsExist(id int) (exist bool) {
	if h.err != nil {
		return
	}
	exist, h.err = todo.IsExist(id)
	return
}

func (h handler) GetData(id int) (result string) {
	if h.err != nil {
		return
	}
	var data Todo
	data, h.err = todo.Get(id)
	if h.err != nil {
		return
	}

}
*/
// Handler is the only one entry point.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var (
		id    int
		err   error
		exist bool
		data  todo.Todo
		b     []byte
	)

	if id, err = strconv.Atoi(request.PathParameters["id"]); err != nil {
		return response.Error(err)
	}
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
