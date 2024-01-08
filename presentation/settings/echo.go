package settings

import (
	"github.com/labstack/echo/v4"
)

func returnAbortWith(ctx *echo.Context, code int, err error) {
	var msg string
	if err != nil {
		msg = err.Error()
	}

	ctx.JSON(code, echo.H{
		"code": code,
		"msg":  msg,
	})
}
