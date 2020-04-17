package models

import (
	"time"
)

type User struct {
	Base
	Email 				string			`gorm:"type:varchar(100);unique_index;" json:"email"`
	Password 			string			`json:"password"`
	LastLogin 			*time.Time		`json:"last_login"`
	VerificationCode 	string			`gorm:"default:null" json:"verification_code"`
	Verified			bool			`gorm:"not null;default:false" json:"verified"`
	Active				bool			`gorm:"not null;default:true" json:"active"`
}
