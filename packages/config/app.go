package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	db *gorm.DB
)

func Connect() {
	fmt.Println("db connecting")

	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatalf("Error loading .env file")
	}

	dbUrl := os.Getenv("DB_URL")
	d, err := gorm.Open("mysql", dbUrl)
	if err != nil {
		logrus.Fatalf("Failure of database connection. error message: %s", err)
	}

	db = d
	fmt.Println("db connected")
}

func GetDB() *gorm.DB {
	return db
}
