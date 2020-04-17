package services

import (
	"github.com/zzlalani/go-users/classes"
	"github.com/zzlalani/go-users/repositories"
)

func GetUsers() []classes.UserResponseGet {
	return repositories.GetUsers()
}

func CreateUser(requestPost *classes.UserRequestPost) {
	repositories.CreateUser(requestPost)
}

func GetUser(id string) classes.UserResponseGet {
	return repositories.GetUser(id)
}

func UpdateUser(id string, requestPut *classes.UserRequestPut) {
	repositories.UpdateUser(id, requestPut)
}

func DeleteUser(id string) {
	repositories.DeleteUser(id)
}