package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:admin@tcp(localhost:3306)/first_demo"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Items{})

	DB = database
}