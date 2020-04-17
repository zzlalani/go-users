package repositories

import (
	"github.com/zzlalani/go-users/classes"
	"github.com/zzlalani/go-users/models"
)

func GetUsers() []classes.UserResponseGet {
	db := models.GetDB()
	var users []models.User
	db.Find(&users)
	models.CloseDB(db)
	var response []classes.UserResponseGet
	for _, user := range users {
		response = append(response, classes.UserResponseGet{
			ID: user.ID,
			Email: user.Email,
		})
	}
	return response
}

func CreateUser(requestPost *classes.UserRequestPost) {
	db := models.GetDB()
	user := models.User{
		Email: requestPost.Email,
		Password: requestPost.Password,
	}
	db.Create(&user)
	db.Model(&user).Update("UpdatedBy", requestPost.Email).Update("CreatedBy", requestPost.Email)
	models.CloseDB(db)
}

func GetUser(id string) classes.UserResponseGet {
	db := models.GetDB()
	user := models.User{}
	db.Where("id = ?", id).First(&user)
	models.CloseDB(db)
	return classes.UserResponseGet{
		ID: user.ID,
		Email: user.Email,
	}
}

func UpdateUser(id string, requestPut *classes.UserRequestPut) {
	db := models.GetDB()
	user := models.User{}
	db.Where("id = ?", id).First(&user)
	user.Email = requestPut.Email
	user.Password = requestPut.Password
	db.Save(&user)
	models.CloseDB(db)
}

func DeleteUser(id string) {
	db := models.GetDB()
	db.Delete(models.User{}, "id = ?", id)
	models.CloseDB(db)
}