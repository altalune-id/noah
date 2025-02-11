package main

import (
	"fmt"
	"net/http"

	"github.com/altalune-id/noah/config"
	"github.com/gin-gonic/gin"
)

func route(cfg *config.Config) *gin.Engine {
	router := gin.Default()
	router.RedirectFixedPath = false

	base := router.Group(fmt.Sprintf("/%s", cfg.Server.Stage))

	base.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Hello darkness my old friend from %s :)", cfg.Server.Stage),
		})
	})

	api := base.Group("/api/v1")
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
