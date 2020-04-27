package controllers

import (
	"github.com/gin-gonic/gin"
	. "github.com/zzlalani/go-users/classes"
	"github.com/zzlalani/go-users/controllers/validations"
	. "github.com/zzlalani/go-users/enumeration"
	"github.com/zzlalani/go-users/services"
	"net/http"
)

func InitUserRoutes(route *gin.Engine) {
	route.GET("/users", getUsers())
	userRoute := route.Group("/user")
	userRoute.POST("/", createUser())
	userRoute.GET("/:id", getUser())
	userRoute.PUT("/:id", updateUser())
	userRoute.DELETE("/:id", deleteUser())
}

func getUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, serviceError := services.GetUsers()
		if (AppError{}) != serviceError {
			c.JSON(serviceError.StatusCode, gin.H{
				"message": serviceError.Error.Error(),
				"code": serviceError.Code,
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, data)
	}
}

func createUser() gin.HandlerFunc {
	return func (c *gin.Context) {
		var requestPost UserRequestCreate
		// Bind
		if err := c.Bind(&requestPost); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
				"code": ErrorCodes["InternalServerError"]["CREATE_USER_DATA_BIND_ERROR"],
			})
			c.Abort()
			return
		}
		// Validate
		if errorMessage := validations.CreateUser(requestPost); errorMessage != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": PrepareValidationErrorsMessages(errorMessage),
				"code": ErrorCodes["UnprocessableEntity"]["CREATE_USER_INPUT_ERROR"],
			})
			c.Abort()
			return
		}
		// Process
		if _, serviceError := services.CreateUser(&requestPost); (AppError{}) != serviceError {
			c.JSON(serviceError.StatusCode, gin.H{
				"message": serviceError.Error.Error(),
				"code": serviceError.Code,
			})
			c.Abort()
			return
		}
		c.Status(http.StatusNoContent)
	}
}

func getUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, serviceError := services.GetUser(c.Param("id"))
		if (AppError{}) != serviceError {
			c.JSON(serviceError.StatusCode, gin.H{
				"message": serviceError.Error.Error(),
				"code": serviceError.Code,
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, data)
	}
}

func updateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestObj UserRequestUpdate
		// Bind
		if err := c.Bind(&requestObj); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
				"code": ErrorCodes["InternalServerError"]["UPDATE_USER_DATA_BIND_ERROR"],
			})
			c.Abort()
			return
		}
		// Validate
		if errorMessage := validations.UpdateUser(requestObj); errorMessage != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": PrepareValidationErrorsMessages(errorMessage),
				"code": ErrorCodes["UnprocessableEntity"]["UPDATE_USER_INPUT_ERROR"],
			})
			c.Abort()
			return
		}
		// Process
		if _, serviceError := services.UpdateUser(c.Param("id"), &requestObj); (AppError{}) != serviceError {
			c.JSON(serviceError.StatusCode, gin.H{
				"message": serviceError.Error.Error(),
				"code": serviceError.Code,
			})
			c.Abort()
			return
		}
		c.Status(http.StatusNoContent)
	}
}

func deleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		if serviceError := services.DeleteUser(c.Param("id")); (AppError{}) != serviceError {
			c.JSON(serviceError.StatusCode, gin.H{
				"message": serviceError.Error.Error(),
				"code": serviceError.Code,
			})
			c.Abort()
			return
		}
 		c.Status(http.StatusNoContent)
	}
}