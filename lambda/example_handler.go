package lambda

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const (
	tableName = "Todos"
	region    = "us-west-2"
)

var (
	sess   *session.Session
	dynamo *dynamodb.DynamoDB
)

// Todo - Todo Item Struct
type Todo struct {
	ID          uint64 `json:"todoId"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Request - Lambda Handler Input
type Request struct {
	Todo Todo `json:"todo"`
}

// Response - Lambda Handler Output
type Response struct {
	Message string `json:"message"`
}

// Init function; called when this file gets loaded / executed
// for the first time.
func init() {
	sess = session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))

	dynamo = dynamodb.New(sess)
}

// HandleRequest - Example Lambda Handler
func HandleRequest(context context.Context, event Request) (*Response, error) {
	log.Println("Starting Handler Execution")
	var err error

	todo := event.Todo
	item, err := dynamodbattribute.MarshalMap(todo)
	if err != nil {
		return nil, err
	}

	_, err = dynamo.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      item,
	})

	if err != nil {
		return nil, err
	}

	message := fmt.Sprintf("Saved Todo: %v", todo.ID)
	return &Response{Message: message}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
