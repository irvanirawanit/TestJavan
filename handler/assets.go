package handler

import (
	"net/http"

	"TestJavan/repository"
	"github.com/labstack/echo/v4"
)

var (
	AssetsRepository = repository.NewAssetsRepository()
)

func AssetsList(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func AssetsDetail(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func AssetsCreate(c echo.Context) error {
	payload := struct {
		Name  string `json:"name"`
		Price int    `json:"price"`
	}{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err := AssetsRepository.AssetsInsert(payload.Name, payload.Price)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "Success")
}

func AssetsUpdate(c echo.Context) error {
	return c.String(http.StatusOK, " Hello, World!")
}

func AssetsDelete(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
