package api

import (
	"net/http"
	"persona/models/user"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Users(c echo.Context) error {
	users := user.FindAll()

	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
	}

	user := user.FindByID(id)

	return c.JSON(http.StatusOK, user)
}
