package api

import (
	"net/http"
	"persona/models/user"
	"strconv"

	"github.com/labstack/echo/v4"
)

type httpError map[string]string

func Users(c echo.Context) error {
	users, err := user.FindAll()
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, httpError{"message": "Error getting users"})
	}

	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpError{"message": "Invalid ID"})
	}

	user, err := user.FindByID(id)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, httpError{"message": "Error getting user"})
	}

	return c.JSON(http.StatusOK, user)
}

func AddUser(c echo.Context) error {
	u := new(user.User)
	err := c.Bind(u)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, httpError{"message": "Form user error bind"})
	}

	users, err := user.Create(*u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

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
	user, err := user.Update(*u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
	}

	err = user.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User deleted"})
}
