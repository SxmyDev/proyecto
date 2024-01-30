package database

import (
	"log"
	"os"
    "fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
    err := loadEnvironmentVariables()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    dsn := getDSN()
    var openErr error
    DB, openErr = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
    if openErr != nil {
        log.Fatalf("Error opening database connection: %v", openErr)
    }

    log.Println("DB connected successfully")
}

func loadEnvironmentVariables() error {
	return godotenv.Load()
}

func getDSN() string {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	return fmt.Sprintf("sqlserver://%s:%s@%s?database=%s", dbUser, dbPassword, dbHost, dbName)
}
