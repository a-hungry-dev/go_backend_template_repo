package api

import (
	"app/model"
	"net/http"

	"gopkg.in/labstack/echo.v4"
)

func GetVersion(c echo.Context) error {
	v := model.NewVersion(1)
	return c.JSON(http.StatusOK, v)
}
