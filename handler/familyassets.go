package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func FamilyAssetsList(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func FamilyAssetsDetail(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func FamilyAssetsCreate(c echo.Context) error {
	return c.String(http.StatusOK, " Hello, World!")
}

func FamilyAssetsUpdate(c echo.Context) error {
	return c.String(http.StatusOK, " Hello, World!")
}

func FamilyAssetsDelete(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
