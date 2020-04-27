package controllers

import (
	"github.com/gin-gonic/gin"
	"time"
)

func InitRoutes() {
	route := gin.Default()
	route.Use(gin.Logger())
	// ping the server
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server pinged AT: " + time.Now().String(),
		})
	})

	InitUserRoutes(route)

	err := route.Run(":8080")
	if err != nil {
		panic(err)
	}
}