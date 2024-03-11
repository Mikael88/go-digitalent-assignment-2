// setup.go
package models

import (
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
)

var DB *gorm.DB

func ConnectDatabase() {
    database, err := gorm.Open(mysql.Open("root:Mikael8898@#@tcp(localhost:3306)/orders_by"), &gorm.Config{})
    if err != nil {
        panic(err)
    }

    err = database.AutoMigrate(&Order{}, &Item{})
    if err != nil {
        panic(err)
    }

    DB = database
}
