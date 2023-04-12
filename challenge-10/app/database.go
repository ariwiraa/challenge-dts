package app

import (
	"challange-10/entity"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "developer"
	password = "babababa"
	dbname   = "middleware_go"
)

func InitializeDatabase() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting database =", err)
	}

	db.Debug().AutoMigrate(&entity.User{}, &entity.Product{})

	fmt.Println("Successfully connected to database")
	return db
}

func GetDB() *gorm.DB {
	var db *gorm.DB

	return db
}
