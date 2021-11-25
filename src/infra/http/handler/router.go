package handler

import (
	"github.com/labstack/echo"
)

func InitRouting(e *echo.Echo, userHandler UserHandler) {
	e.POST("/user/create", userHandler.Create())
	e.GET("/user/get", userHandler.Get())
	e.PUT("/user/update", userHandler.Update())
}
