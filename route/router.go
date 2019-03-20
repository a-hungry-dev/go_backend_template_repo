package route

import (
	"app/api"

	"gopkg.in/labstack/echo.v4"
)

func Init() *echo.Echo {

	e := echo.New()

	// Set Bundle MiddleWare
	// e.Use(middleware.Logger())
	// e.Use(middleware.Gzip())
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	// }))

	// Routes
	e.GET("/api/version", api.GetVersion)

	return e
}
