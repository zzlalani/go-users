package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zzlalani/go-users/classes"
	"github.com/zzlalani/go-users/services"
	"net/http"
)

func InitUserRoutes(route *gin.Engine) {
	route.GET("/users", getUsers())
	route.POST("/user", createUser())
	route.GET("/user/:id", getUser())
	route.PUT("/user/:id", updateUser())
	route.DELETE("/user/:id", deleteUser())
}

func getUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, services.GetUsers())
	}
}

func createUser() gin.HandlerFunc {
	return func (c *gin.Context) {
		requestPost := &classes.UserRequestPost{}
		c.Bind(&requestPost)
		services.CreateUser(requestPost)
		c.Status(http.StatusNoContent)
	}
}

func getUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, services.GetUser(c.Param("id")))
	}
}

func updateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestPut := &classes.UserRequestPut{}
		c.Bind(&requestPut)
		services.UpdateUser(c.Param("id"), requestPut)
		c.Status(http.StatusNoContent)
	}
}

func deleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		services.DeleteUser(c.Param("id"))
		c.Status(http.StatusNoContent)
	}
}