package db

import (
	"root/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	url := config.Config("URL_DB")
	if url == "nil" {
		panic("Database url not found")
	}
	var err error
	DB, err = gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		panic("Could not connect to database")
	}

	// err = DB.AutoMigrate(&model.Post{})
	// if err != nil {
	// 	panic("Could not auto migrate database")
	// }

	// fmt.Println("Database Migrated")
}
