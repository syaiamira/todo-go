package database

import (
	"log"
	"todo-cognixus/config"
	"todo-cognixus/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectSQLite() {
	var err error

	DB, err = gorm.Open(sqlite.Open(config.DATABASE_URL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = DB.AutoMigrate(&model.User{}, &model.Todo{})
	if err != nil {
		log.Fatal(err)
	}
}
