package api

import (
	"net/http"
	"persona/models/account"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Accounts(c echo.Context) error {
	accounts, err := account.FindAll()
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, httpError{"message": "Error getting accounts"})
	}

	return c.JSON(http.StatusOK, accounts)
}

func GetAccount(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpError{"message": "Invalid ID"})
	}

	account, err := account.FindByID(id)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, httpError{"message": "Error getting account"})
	}

	return c.JSON(http.StatusOK, account)
}

func AddAccount(c echo.Context) error {
	a := new(account.Account)
	err := c.Bind(a)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, httpError{"message": "Form account error bind"})
	}

	accounts, err := account.Create(*a)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, accounts)
}

func UpdateAccount(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
	}

	a := new(account.Account)
	err = c.Bind(a)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, httpError{"message": "Form account error bind"})
	}

	a.ID = id
	account, err := account.Update(*a)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, account)
}

func DeleteAccount(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
	}

	err = account.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, httpError{"message": "Account deleted"})
}
