package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitMySQL() (*gorm.DB, error) {
	// set environtment variable for setup mysql database
	DB_USER := os.Getenv("DB_USER")
	if DB_USER == "" {
		log.Println("Environment variable DB_USER must be set")
	}

	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	if DB_PASSWORD == "" {
		log.Println("Environment variable DB_PASSWORD must be set")
	}

	DB_HOST := os.Getenv("DB_HOST")
	if DB_HOST == "" {
		log.Println("Environment variable DB_HOST must be set")
	}

	DB_PORT := os.Getenv("DB_PORT")
	if DB_PORT == "" {
		log.Println("Environment variable DB_PORT must be set")
	}

	DB_NAME := os.Getenv("DB_NAME")
	if DB_NAME == "" {
		log.Println("Environment variable DB_NAME must be set")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InitPostgreSQL() (*gorm.DB, error) {
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", "db", "5432", "postgres", "password", "sawer", "disable")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
