package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

// @Reference: https://medium.com/@the.hasham.ali/how-to-use-uuid-key-type-with-gorm-cc00d4ec7100
// Base contains common columns for all tables.
type Base struct {
	ID        	uuid.UUID 	`gorm:"type:uuid;primary_key;"`
	CreatedAt 	time.Time
	CreatedBy	string
	UpdatedAt 	time.Time
	UpdatedBy	string
	DeletedAt 	*time.Time 	`sql:"index"`
	DeletedBy	string		`gorm:"default:null"`
}
