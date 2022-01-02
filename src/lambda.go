package main

import (
  "time"
  "encoding/json"
  "context"

  "github.com/aws/aws-lambda-go/lambda"
  "github.com/aws/aws-lambda-go/events"

  "github.com/sean-callahan/mcquery"
)

type Response struct {
  StatusCode int               `json:"statusCode"`
  Headers    map[string]string `json:"headers"`
  Body       map[string]map[string]string            `json:"body"`
}

func LambdaHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	mcq, err := mcquery.Dial("mc.backman.fyi:25565", time.Second)
	if err != nil {
		panic(err)
	}

	status, _, err := mcq.GetStatus()
	if err != nil {
		panic(err)
	}

  statusJson, err := json.Marshal(status)
  return events.APIGatewayProxyResponse{
    Body: string(statusJson), 
    StatusCode: 200, 
    Headers: map[string]string{"Content-Type": "application/json"},
  }, nil
}
 
func main() {
  lambda.Start(LambdaHandler)
}
