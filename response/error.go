package response

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

// Error returns 500 error response.
func Error(err error) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("%+v\n", err)
	return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Internal Server Error!"}, nil
}
