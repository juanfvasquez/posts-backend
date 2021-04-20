package db

import (
	"gorm.io/gorm"

	"postsbackend/models"
)

var db *gorm.DB

func GetConnection() *gorm.DB {
	if db != nil {
		return db
	}
	db = getDBInstance()
	return db
}

func Migrate() {
	_ = GetConnection()
	db.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Comment{},
		&models.Reaction{},
		&models.Group{},
	)
}
