package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"user_wallet/config"
	"user_wallet/struct"
)

var DB *gorm.DB

func Connect() {
	var err error

	config.LoadConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
		config.AppConfig.Database.Host,
		config.AppConfig.Database.User,
		config.AppConfig.Database.Password,
		config.AppConfig.Database.DbName,
		config.AppConfig.Database.SSLMode,
	)

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
