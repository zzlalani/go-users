package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	uuid "github.com/satori/go.uuid"
)

var dialect = "postgres"
var connection = "host=localhost port=5432 user=postgres dbname=go-user password=root sslmode=disable"

func GetDB() *gorm.DB {
	db, err := gorm.Open(
		dialect,
		connection,
	)
	if err != nil {
		panic(err)
	}
	return db
}

func CloseDB(db *gorm.DB) {
	defer db.Close()
}

func InitDB() {
	db, err := gorm.Open(
		dialect,
		connection,
	)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	initMigrate(db)
}

func initMigrate (db *gorm.DB) {
	db.LogMode(true)
	db.AutoMigrate(
		&User{},
	)
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	_uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", _uuid)
}