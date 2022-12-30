package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func FamilyList(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func FamilyDetail(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func FamilyCreate(c echo.Context) error {
	return c.String(http.StatusOK, " Hello, World!")
}

func FamilyUpdate(c echo.Context) error {
	return c.String(http.StatusOK, " Hello, World!")
}

func FamilyDelete(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
