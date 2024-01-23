package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetHealthHandler(e echo.Context) error {
	return e.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}
