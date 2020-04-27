package validations

import (
	. "github.com/zzlalani/go-users/classes"
	"gopkg.in/bluesuncorp/validator.v9"
)

func CreateUser(requestObj UserRequestCreate) []string {
	validate := validator.New()
	var errorMessage []string
	if err := validate.Struct(requestObj); err != nil {
		for _, fieldErr := range err.(validator.ValidationErrors) {
			if fieldErr.Field() == "Email" && fieldErr.Tag() == "required" {
				errorMessage = append(errorMessage, "Email address is required")
			}
			if fieldErr.Field() == "Email" && fieldErr.Tag() == "email" {
				errorMessage = append(errorMessage, "Invalid email address")
			}
			if fieldErr.Field() == "Password" && fieldErr.Tag() == "required" {
				errorMessage = append(errorMessage, "Password is required")
			}
			if fieldErr.Field() == "Password" && (fieldErr.Tag() == "min" || fieldErr.Tag() == "max") {
				errorMessage = append(errorMessage, "Password must be between 8 and 32 characters")
			}
		}
		return errorMessage
	}
	return nil
}

func UpdateUser(requestObj UserRequestUpdate) []string {
	validate := validator.New()
	var errorMessage []string
	if err := validate.Struct(requestObj); err != nil {
		for _, fieldErr := range err.(validator.ValidationErrors) {
			if fieldErr.Field() == "Email" && fieldErr.Tag() == "required" {
				errorMessage = append(errorMessage, "Email address is required")
			}
			if fieldErr.Field() == "Email" && fieldErr.Tag() == "email" {
				errorMessage = append(errorMessage, "Invalid email address")
			}
			if fieldErr.Field() == "Password" && fieldErr.Tag() == "required" {
				errorMessage = append(errorMessage, "Password is required")
			}
			if fieldErr.Field() == "Password" && (fieldErr.Tag() == "min" || fieldErr.Tag() == "max") {
				errorMessage = append(errorMessage, "Password must be between 8 and 32 characters")
			}
		}
		return errorMessage
	}
	return nil
}