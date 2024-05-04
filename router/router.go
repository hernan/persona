package router

import (
	"net/http"
	"os"
	"persona/api"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	v1 := e.Group("/api/v1")

	users := v1.Group("/users", apiTokenMiddleware)
	users.GET("", api.Users)
	users.GET("/:id", api.GetUser)
	users.POST("", api.AddUser)
	users.PUT("/:id", api.UpdateUser)
	users.DELETE("/:id", api.DeleteUser)

	accounts := v1.Group("/accounts")
	accounts.GET("", api.Accounts)
	accounts.GET("/:id", api.GetAccount)
	accounts.POST("", api.AddAccount)
	accounts.PUT("/:id", api.UpdateAccount)
	accounts.DELETE("/:id", api.DeleteAccount)

	e.Logger.Fatal(e.Start(":8085"))
}

func apiTokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := os.Getenv("TOKEN")
		if token == "" {
			return echo.NewHTTPError(http.StatusInternalServerError, "Token not defined")
		}

		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization header")
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authorization header")
		}

		providedToken := strings.TrimPrefix(authHeader, "Bearer ")
		if providedToken != token {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		return next(c)
	}
}
