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
	defer config.DisconnectDB(db)

	routes.Routes()
}
