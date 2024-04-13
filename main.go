package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		data := map[string]string{"message": "Hello, World!"}
		return c.JSON(http.StatusOK, data)

		// return c.String(http.StatusOK, "Hello, World!")
		// return c.HTML(http.StatusOK, "<h1>Hello, World!</h1>")
	})

	e.GET("/users", func(c echo.Context) error {
		return c.String(http.StatusOK, "Users")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
