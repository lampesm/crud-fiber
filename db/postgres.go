package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/lampesm/crud-fiber/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	env := godotenv.Load(".env")

	if env != nil {
		panic("Failed to load .env file")
	}

	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbName := os.Getenv("POSTGRES_DB")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran", dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to postgres database")
	}
	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		panic("Failed: Unable to migrate your postgres database")
	}

	return db
}

func Close(db *gorm.DB) {
	dbPsql, err := db.DB()
	if err != nil {
		panic("Failed: postgres database connection")
	}
	err = dbPsql.Close()
	if err != nil {
		panic("FAiled: unable to close postgres connection database")
	}
}
