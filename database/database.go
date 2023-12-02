package database

import (
	"log"
	"todo-cognixus/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectSQLite() {
	var err error

	DB, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = DB.AutoMigrate(&model.Todo{})
	if err != nil {
		log.Fatal(err)
	}

}
