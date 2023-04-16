package app

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "developer"
	password = "babababa"
	dbname   = "hacktiv_go"
)

func InitializeDatabase() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)

	// db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal("error connecting database =", err)
	// }

	// db.Debug().AutoMigrate(&entity.Book{})
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")
	return db
}
