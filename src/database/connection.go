package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _DbConnect = "host=localhost user=postgres password=* dbname=gorm port=5432"
var DB *gorm.DB

func DbConnection() {

	var error error

	DB, error = gorm.Open(postgres.Open(_DbConnect), &gorm.Config{})

	if error != nil {
		log.Fatalln(error)
	} else {
		log.Println("Db Connect")
	}

}
