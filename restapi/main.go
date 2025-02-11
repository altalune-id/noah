package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/altalune-id/noah/config"
	adapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Order struct {
	ID      string `json:"id"`
	Product string `json:"product"`
	Qty     int    `json:"qty"`
}

var (
	cfg       *config.Config
	ginLambda *adapter.GinLambdaV2
	router    *gin.Engine
)

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
	{
		ID:      "3",
		Product: "Monitor",
		Qty:     1,
	},
}

func isRunningInLambda() bool {
	_, exists := os.LookupEnv("AWS_LAMBDA_FUNCTION_NAME")
	return exists
}

func handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	// warm start
	return ginLambda.ProxyWithContext(ctx, request)
}

func init() {
	// cold start: load config once
	var err error
	cfg, err = config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	router = route(cfg)
	ginLambda = adapter.NewV2(router)
}

func main() {
	if !isRunningInLambda() {
		router.Run(fmt.Sprintf(":%d", cfg.Server.Port))
		return
	}

	// cold start
	lambda.Start(handler)
}
