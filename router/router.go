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
	v1.Use(apiTokenMiddleware)

	v1.GET("/users", api.Users)
	v1.GET("/users/:id", api.GetUser)
	v1.POST("/users", api.AddUser)
	v1.PUT("/users/:id", api.UpdateUser)
	v1.DELETE("/users/:id", api.DeleteUser)

	v1Acc := v1.Group("/accounts")
	v1Acc.GET("", api.Accounts)
	v1Acc.GET("/:id", api.GetAccount)
	v1Acc.POST("", api.AddAccount)
	v1Acc.PUT("/:id", api.UpdateAccount)
	v1Acc.DELETE("/:id", api.DeleteAccount)

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
