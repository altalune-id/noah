package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/altalune-id/noah/awscli"
	"github.com/altalune-id/noah/config"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
)

var (
	cfg         *config.Config
	apigwClient *awscli.ApiGatewayClient
)

var connections = make(map[string]struct{})

func isRunningInLambda() bool {
	_, exists := os.LookupEnv("AWS_LAMBDA_FUNCTION_NAME")
	return exists
}

func handler(ctx context.Context, req events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	// warm start
	switch req.RequestContext.RouteKey {
	case "$connect":
		connections[req.RequestContext.ConnectionID] = struct{}{}
		log.Printf("Connected: %s", req.RequestContext.ConnectionID)

	case "$disconnect":
		delete(connections, req.RequestContext.ConnectionID)
		log.Printf("Disconnected: %s", req.RequestContext.ConnectionID)

	case "$default":
		var data map[string]string
		json.Unmarshal([]byte(req.Body), &data)

		message := data["message"]
		for connID := range connections {
			_, err := apigwClient.PostToConnection(&apigatewaymanagementapi.PostToConnectionInput{
				ConnectionId: awssdk.String(connID),
				Data:         []byte(message),
			})
			if err != nil {
				log.Printf("Error sending message: %v", err)
			}
		}
	}

	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

func init() {
	// cold start: load config once
	var err error
	cfg, err = config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	apigwClient = awscli.NewApiGatewayClient(cfg)
}

func main() {
	if !isRunningInLambda() {
		log.Fatalln("This should only run in Lambda")
		return
	}

	// cold start
	lambda.Start(handler)
}
