package app

import (
	"fmt"
	"gin-gorm/entity"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "developer"
	password = "babababa"
	dbname   = "hacktiv_go"
)

func InitializeDatabase() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting database =", err)
	}

	db.Debug().AutoMigrate(&entity.Book{})

	fmt.Println("Successfully connected to database")
	return db
}
