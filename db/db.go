package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"user_wallet/struct"
)

var DB *gorm.DB

func Connect() {
	var err error
	dsn := "host=localhost user=postgres password=Martin1992& dbname=go_test sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	err = DB.AutoMigrate(&structs.User{}, &structs.Wallet{})
	if err != nil {
		log.Fatal("failed to migrate database:", err)
	}

	log.Println("Connected to the database successfully and schema migrated!")
}
