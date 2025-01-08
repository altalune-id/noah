package main

import (
	"context"
	"fmt"
	"log"

	"github.com/altalune-id/noah/config"
	adapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Order struct {
	ID      string `json:"id"`
	Product string `json:"product"`
	Qty     int    `json:"qty"`
}

var ginLambda *adapter.GinLambdaV2
var orders = []Order{
	{
		ID:      "1",
		Product: "Laptop",
		Qty:     1,
	},
	{
		ID:      "2",
		Product: "Mouse",
		Qty:     2,
	},
}

func lambdaProxy(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return ginLambda.ProxyWithContext(ctx, request)
}

func main() {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	if cfg.Server.Mode == "local" {
		router := route(cfg)
		router.Run(fmt.Sprintf(":%d", cfg.Server.Port))
		return
	}

	if ginLambda == nil {
		router := route(cfg)
		ginLambda = adapter.NewV2(router)
	}
	lambda.Start(lambdaProxy)
}
