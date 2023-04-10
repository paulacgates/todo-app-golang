package database

import (
	oracle "github.com/godoes/gorm-oracle"
	"github.com/paulacgates/todo-app-golang/models"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func InitDatabase() {
	url := ""
	DB, err = gorm.Open(oracle.Open(url), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&models.Todo{})
}
