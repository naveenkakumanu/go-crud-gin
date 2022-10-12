package db

import (
	"fmt"

	"github.com/naveenkakumanu/go-crud-gin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

// Database Connection
func (database *DB) InitDB(configuration models.DBConfig) {
	config := configuration.DB
	dsn := config.UserName + ":" + config.Credential + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	database.db = db
}

func (database *DB) GetDB() *gorm.DB {
	return database.db
}
