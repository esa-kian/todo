package main

import (
	"todo/src/config"
	"todo/src/routes"

	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()
)

func main() {
	type Todo struct {
		gorm.Model
		Name        string
		Description string
	}
	db.AutoMigrate(&Todo{})

	defer config.DisconnectDB(db)

	routes.Routes()
}
