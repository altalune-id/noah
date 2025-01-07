package main

import (
	"context"
	"net/http"

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

func init() {
	router := gin.Default()
	router.RedirectFixedPath = false

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello darkness my old friend DEV",
		})
	})

	api := router.Group("/api/v1")
	{
		order := api.Group("/orders")
		order.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, orders)
		})
		order.POST("", func(c *gin.Context) {
			var order Order
			if err := c.BindJSON(&order); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			orders = append(orders, order)
			c.JSON(http.StatusOK, order)
		})
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Resource not found",
		})
	})

	ginLambda = adapter.NewV2(router)
}

func handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return ginLambda.ProxyWithContext(ctx, request)
}

func main() {
	lambda.Start(handler)
}
