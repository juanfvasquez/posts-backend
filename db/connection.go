package db

import (
	"log"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"../models"
)

var db *gorm.DB
var err error

const (
	user     = "postgres"
	password = "123456"
	dbname   = "posts"
	host     = "localhost"
	port     = "5432"
)

func GetConnection() *gorm.DB {
	if db != nil {
		return db
	}
	conexion := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
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
