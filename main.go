package main

import (
	"persona/models"
	"persona/router"
)

func main() {
	models.DBInit()
	router.Init()
}
