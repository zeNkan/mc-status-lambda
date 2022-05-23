package main

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/sean-callahan/mcquery"
)

func LambdaHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var mcq, err = mcquery.Dial(os.Getenv("MC_SERVER_ADDR")+":"+os.Getenv("MC_SERVER_PORT"), time.Second)
	if err != nil {
		panic(err)
	}

	statusMap, _, err := mcq.GetStatus()
	if err != nil {
		panic(err)
	}

	statusMap["hostname"] = os.Getenv("MC_SERVER_ADDR")
	delete(statusMap, "hostip")

	statusJson, err := json.Marshal(statusMap)
	return events.APIGatewayProxyResponse{
		Body:       string(statusJson),
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
	}, nil
}

func main() {
	lambda.Start(LambdaHandler)
}
