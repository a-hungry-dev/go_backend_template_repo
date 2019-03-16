package api

import (
	"app/model"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

func GetVersion(c echo.Context) error {
	v := model.NewVersion(1)
	return c.JSON(fasthttp.StatusOK, v)
}