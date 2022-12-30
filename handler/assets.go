package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AssetsList(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func AssetsDetail(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func AssetsCreate(c echo.Context) error {
	return c.String(http.StatusOK, " Hello, World!")
}

func AssetsUpdate(c echo.Context) error {
	return c.String(http.StatusOK, " Hello, World!")
}

func AssetsDelete(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
