package response

import (
	"github.com/aws/aws-lambda-go/events"
)

// Success returns 200 ok response.
func Success(body string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: body}, nil
}
