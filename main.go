package main

import (
	"github.com/zzlalani/go-users/controllers"
	"github.com/zzlalani/go-users/models"
)

func main () {
	models.InitDB()
	controllers.InitRoutes()
}