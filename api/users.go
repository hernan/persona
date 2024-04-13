package api

import (
	"net/http"
	"persona/models"

	"github.com/labstack/echo/v4"
)

func Users(c echo.Context) error {
	users := models.UserDB

	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
