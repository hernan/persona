package main

import (
	"persona/models"
	"persona/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	models.DBInit()
	router.Init()

	defer models.MyDb.Close()
}
