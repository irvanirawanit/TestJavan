package main

import (
	"net/http"

	"TestJavan/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "TestJavan")
	})

	// family
	e.GET("/family", handler.FamilyList)
	e.GET("/family/:id", handler.FamilyDetail)
	e.POST("/family", handler.FamilyCreate)
	e.PUT("/family/:id", handler.FamilyUpdate)
	e.DELETE("/family/:id", handler.FamilyDelete)

	// assets
	e.GET("/assets", handler.AssetsList)
	e.GET("/assets/:id", handler.AssetsDetail)
	e.POST("/assets", handler.AssetsCreate)
	e.PUT("/assets/:id", handler.AssetsUpdate)
	e.DELETE("/assets/:id", handler.AssetsDelete)

	// family_assets
	e.GET("/family_assets", handler.FamilyAssetsList)
	e.GET("/family_assets/:id", handler.FamilyAssetsDetail)
	e.POST("/family_assets", handler.FamilyAssetsCreate)
	e.PUT("/family_assets/:id", handler.FamilyAssetsUpdate)
	e.DELETE("/family_assets/:id", handler.FamilyAssetsDelete)

	e.Logger.Fatal(e.Start(":1323"))
}
