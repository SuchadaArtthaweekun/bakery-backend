package database

import (
	"fmt"
	"go-fiber/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	dsn := "host=172.21.192.1 user=postgres password=password dbname=my_db port=5432 sslmode=disable"
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Could not connect database")
		fmt.Println(err.Error())
	} else {
		fmt.Println("Database connected")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
	connection.AutoMigrate(&models.Product{})

}
