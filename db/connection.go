package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func (database *DB) InitDB() {

	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	database.db = db
}

func (database *DB) GetDB() *gorm.DB {
	return database.db
}
