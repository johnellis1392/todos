package lambda

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

var (
	ERRNameNotProvided = errors.New("No name was provided in the HTTP body")
)

/**
 * Example Lambda Handler for AWS API Gateway Lambda Proxy Integration
 */
func HandleRequest_ApiGateway(context context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Processing Lambda Request %s\n", request.RequestContext.RequestID)

	if len(request.Body) < 1 {
		return events.APIGatewayProxyResponse{}, ERRNameNotProvided
	}

	return events.APIGatewayProxyResponse{
		Body:       "Hello " + request.Body,
		StatusCode: 200,
	}, nil
}

// func main() {
// 	lambda.Start(HandleRequest)
// }
