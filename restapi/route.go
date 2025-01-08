package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func route() *gin.Engine {
	router := gin.Default()
	router.RedirectFixedPath = false

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello darkness my old friend :)",
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

	return router
}
