package api

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"persona/models/account"
	"persona/models/session"
	"strings"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	acc := account.Account{}
	err := c.Bind(&acc)
	if err != nil {
		return echo.ErrBadRequest
	}

	if (acc.Name == nil || strings.TrimSpace(*acc.Name) == "") || (acc.Password == nil || strings.TrimSpace(*acc.Password) == "") {
		return echo.ErrBadRequest
	}

	accountExists := checkAccountExits(*acc.Name, *acc.Password)
	if !accountExists {
		return echo.ErrUnauthorized
	}

	token, err := generateRandomToken()
	if err != nil {
		return echo.ErrInternalServerError
	}

	storeToken(token)

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})

	// The token will be invalidated when the user logs out
	// The token will be stored in the server and checked for each request
	// The token will be stored in the server and invalidated when the user logs out
	// The token will be stored in the server and invalidated after a certain period of time
}

func storeToken(token string) {
	session.Create(session.Session{
		UserID:  1,
		Session: &token,
	})
}

func generateRandomToken() (string, error) {
	token := make([]byte, 16)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(token), nil
}

func checkAccountExits(name string, password string) bool {
	acc, err := account.FindByName(name)
	if err != nil {
		return false
	}

	if *acc.Password != password {
		return false
	}

	return acc.ID != 0
}

func Logout(c echo.Context) error {
	return nil
}
