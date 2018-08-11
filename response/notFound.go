package response

import (
	"github.com/aws/aws-lambda-go/events"
)

// NotFound returns 404 not found response.
func NotFound(body string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{StatusCode: 404, Body: body}, nil
}
