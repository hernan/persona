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

func AddUser(c echo.Context) error {
	u := new(user.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	users := user.Create(*u)

	return c.JSON(http.StatusCreated, users)
}

func UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
	}

	u := new(user.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	u.ID = id
	user := user.Update(*u)

	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
	}

	user.Delete(id)

	return c.JSON(http.StatusOK, map[string]string{"message": "User deleted"})
}
