package handler

import (
	"net/http"

	"TestJavan/repository"
	"github.com/labstack/echo/v4"
)

var (
	FamilyRepository = repository.NewFamilyRepository()
)

func FamilyList(c echo.Context) error {
	data, err := FamilyRepository.FamilyList()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, data)
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
