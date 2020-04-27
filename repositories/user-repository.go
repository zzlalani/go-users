package repositories

import (
	"errors"
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	. "github.com/zzlalani/go-users/classes"
	. "github.com/zzlalani/go-users/enumeration"
	. "github.com/zzlalani/go-users/models"
	"github.com/zzlalani/go-users/utilities"
	"log"
)

func GetUsers() ([]UserResponseGet, AppError) {
	db := GetDB()
	var users []User
	var response []UserResponseGet
	if status := db.Find(&users); status.Error != nil {
		log.Println(status.Error)
		return response, InternalServerError(
			ErrorCodes["InternalServerError"]["GET_USERS_UNKNOWN_ERROR"],
			errors.New("unknown error, please try again"),
		)
	}
	for _, user := range users {
		response = append(response, UserResponseGet{
			ID: user.ID,
			Email: user.Email,
		})
	}
	CloseDB(db)
	return response, AppError{}
}

func CreateUser(requestPost *UserRequestCreate) (uuid.UUID, AppError) {
	db := GetDB()
	user := User{
		Email: requestPost.Email,
		Password: utilities.HashText(requestPost.Password),
		Base: Base{
			UpdatedBy: requestPost.Email,
			CreatedBy: requestPost.Email,
		},
	}
	if status := db.Create(&user); status.Error != nil {
		log.Println(status.Error)
		switch status.Error.(*pq.Error).Code.Name() {
		case "unique_violation":
			return uuid.Nil, Conflict(
				ErrorCodes["Conflict"]["CREATE_USER_ALREADY_EXISTS"],
				errors.New("email already exists"),
			)
		default:
			return uuid.Nil, InternalServerError(
				ErrorCodes["InternalServerError"]["CREATE_USER_UNKNOWN_ERROR"],
				errors.New("unknown error, please try again"),
			)
		}
	}
	CloseDB(db)
	return user.ID, AppError{}
}

func GetUser(id string) (UserResponseGet, AppError) {
	db := GetDB()
	user := User{}
	if status := db.Where("id = ?", id).First(&user); status.Error != nil {
		log.Println(status.Error)
		return UserResponseGet{}, InternalServerError(
			ErrorCodes["InternalServerError"]["GET_USER_UNKNOWN_ERROR"],
			errors.New("unknown error, please try again"),
		)
	}
	CloseDB(db)
	return UserResponseGet{
		ID: user.ID,
		Email: user.Email,
	}, AppError{}
}

func UpdateUser(id string, requestObj *UserRequestUpdate) (bool, AppError) {
	db := GetDB()
	user := User{}
	modelObj := db.Where("id = ?", id).First(&user)
	if modelObj.RecordNotFound() {
		return false, NotFound(
			ErrorCodes["NotFound"]["UPDATE_USER_NOT_FOUND"],
			errors.New("invalid user id"),
		)
	}
	user.Email = requestObj.Email
	user.Password = utilities.HashText(requestObj.Password)
	user.UpdatedBy = user.Email
	status := db.Save(&user)
	if status.Error != nil {
		log.Println(status.Error)
		switch status.Error.(*pq.Error).Code.Name() {
		case "unique_violation":
			return false, Conflict(
				ErrorCodes["Conflict"]["UPDATE_USER_ALREADY_EXISTS"],
				errors.New("email already exists"),
			)
		default:
			return false, InternalServerError(
				ErrorCodes["InternalServerError"]["UPDATE_USER_UNKNOWN_ERROR"],
				errors.New("unknown error, please try again"),
			)
		}
	}
	CloseDB(db)
	return true, AppError{}
}

func DeleteUser(id string) AppError {
	db := GetDB()
	if status := db.Delete(User{}, "id = ?", id); status.Error != nil {
		log.Println(status.Error)
		return InternalServerError(
			ErrorCodes["InternalServerError"]["DELETE_USER_UNKNOWN_ERROR"],
			errors.New("unknown error, please try again"),
		)
	}
	CloseDB(db)
	return AppError{}
}