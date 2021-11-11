package utils

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	dsn := "host=172.31.48.1 user=postgres password=39053372 dbname=postgres port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if(err != nil) {
		panic("Fail to connect to DB")
	}else {
		fmt.Println("Connected to Database")
	}

	DB = database
	
}