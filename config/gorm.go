package config

import (
	"assignment2/httpserver/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDb() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/assignment2?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.Order{}, &models.Item{})
	return db, err
}
