package config

import (
	"bitespeedtask/models"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDB() {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")

	URL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS,HOST, DBNAME)
	db, err := gorm.Open(mysql.Open(URL))
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Initialized DB successfully!!")

	err = db.AutoMigrate(&models.Contact{})
	if err != nil {
		return
	}

	DB = db
}
