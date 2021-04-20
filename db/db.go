package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func getDBInstance() (db *gorm.DB) {
	dbConn := os.Getenv("DB_CONNECTION")
	var connStr string
	var err error
	switch dbConn {
	case "postgres":
		connStr = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_DATABASE"),
		)
		db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	case "mysql":
		connStr = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_DATABASE"),
		)
		db, err = gorm.Open(mysql.Open(connStr), &gorm.Config{})
	default:
		db, err = gorm.Open(sqlite.Open(os.Getenv("DB_DATABASE")), &gorm.Config{})
	}
	if err != nil {
		log.Println("DB connection error...")
		panic(err)
	}
	return
}
