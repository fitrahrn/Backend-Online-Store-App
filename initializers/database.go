package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	host := os.Getenv("HOST")
	// user := os.Getenv("USER")
	// pass := os.Getenv("PASSWORD")
	// port := os.Getenv("PORT")
	// name := os.Getenv("NAME")
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", host, "postgres", "postgres", "online_stores", 5432)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}
