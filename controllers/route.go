package controllers

import (
	"github.com/gin-gonic/gin"
	"time"
)

var route *gin.Engine

func InitRoutes() {
	route = gin.Default()
	// ping the server
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server pinged AT: " + time.Now().String(),
		})
	})

	InitUserRoutes(route)

	err := route.Run()
	if err != nil {
		panic(err)
	}
}