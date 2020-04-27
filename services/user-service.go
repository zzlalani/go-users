package services

import (
	uuid "github.com/satori/go.uuid"
	. "github.com/zzlalani/go-users/classes"
	. "github.com/zzlalani/go-users/enumeration"
	"github.com/zzlalani/go-users/repositories"
)

func GetUsers() ([]UserResponseGet, AppError) {
	return repositories.GetUsers()
}

func CreateUser(requestPost *UserRequestCreate) (uuid.UUID, AppError) {
	return repositories.CreateUser(requestPost)
}

func GetUser(id string) (UserResponseGet, AppError) {
	return repositories.GetUser(id)
}

func UpdateUser(id string, requestObj *UserRequestUpdate) (bool, AppError) {
	return repositories.UpdateUser(id, requestObj)
}

func DeleteUser(id string) AppError {
	return repositories.DeleteUser(id)
}