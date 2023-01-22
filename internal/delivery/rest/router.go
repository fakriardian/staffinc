package rest

import "github.com/labstack/echo/v4"

func LoadRoutes(e *echo.Echo, handler *handler) {

	apiGroup := e.Group("/api")
	apiGroup.POST("/input-harga", handler.InputHarga)
	apiGroup.GET("/check-harga", handler.CheckHarga)

	apiGroup.POST("/saldo", handler.CheckSaldo)
}
