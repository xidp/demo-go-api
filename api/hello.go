package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetHelloHandler(e echo.Context) error {
	return e.JSON(http.StatusOK, map[string]string{
		"hello": "Universe",
	})
}
