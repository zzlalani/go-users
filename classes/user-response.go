package classes

import (
	uuid "github.com/satori/go.uuid"
)

type UserResponseGet struct {
	ID			uuid.UUID	`json:"id"`
	Email		string	`json:"email"`
}