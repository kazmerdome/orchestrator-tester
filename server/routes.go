package server

import (
	"github.com/labstack/echo"
	"gitlab.com/kazmerdome/orchestrator-tester/controller"
)

// GetRoutes ...
func GetRoutes(e *echo.Echo) {
	info := new(controller.InfoController)
	e.GET("*", info.GetInfo)
	e.POST("*", info.GetInfo)
	e.PUT("*", info.GetInfo)
	e.PATCH("*", info.GetInfo)
	e.DELETE("*", info.GetInfo)
	e.CONNECT("*", info.GetInfo)
	e.OPTIONS("*", info.GetInfo)
	e.TRACE("*", info.GetInfo)
}
