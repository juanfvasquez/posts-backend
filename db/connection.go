package db

import (
	"log"
	"fmt"
	"os"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"

	"../models"
)

var db *gorm.DB
var err error

func GetConnection() *gorm.DB {
	if db != nil {
		return db
	}
	conexion := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
	)
	db, err = gorm.Open(postgres.Open(conexion), &gorm.Config{})
	if err != nil {
		log.Println("Error en la conexión...")
		panic(err)
	}
	return db
}

func Migrate() {
	_ = GetConnection()
	db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
}
