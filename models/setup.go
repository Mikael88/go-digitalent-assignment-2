package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    dsn := "root:Mikael8898@#@tcp(localhost:3306)/orders_by?parseTime=True"
    database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database!")
    }

    database.AutoMigrate(&Order{}, &Item{})

    DB = database
}
