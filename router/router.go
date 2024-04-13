package router

import (
	"persona/api"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	v1 := e.Group("/api/v1")
	v1.GET("/users", api.Users)
	v1.GET("/users/:id", api.GetUser)
	v1.GET("/members", api.Members)

	e.Logger.Fatal(e.Start(":8085"))
}
