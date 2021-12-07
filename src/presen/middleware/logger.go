package middleware

import (
	"github.com/labstack/echo"

	"techtrain-mission/src/core/logger"
)

func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request()
		logger.HttpLogging("success", r)
		if err := next(c); err != nil {
			c.Error(err)
			return err
		}
		return nil
	}
}
