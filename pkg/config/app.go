package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var (
	db * gorm.DB
)

func Connect() {
    pwd, err := os.Getwd()
    if err != nil {
        panic(err)
    }
    fmt.Println("Current working directory:", pwd)

    envPath := filepath.Join(pwd, "..", "..", ".env") 
    fmt.Println("Attempting to load .env file from:", envPath)
    
    err = godotenv.Load(envPath)
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    DB_Link := os.Getenv("DB_URL")
    fmt.Println("DB URL:", DB_Link)
    d, err := gorm.Open("mysql", DB_Link)
    if err != nil {
        panic(fmt.Sprintf("Failed to connect to database: %v", err))
    }
    db = d
}

func GetDB() *gorm.DB{
	return db
}