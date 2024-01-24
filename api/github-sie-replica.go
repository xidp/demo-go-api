package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetGithubHandler(e echo.Context) error {
	return e.JSON(http.StatusOK, map[string]string{
		"version":     "v0.1.0",
		"gitSHA":      "87998bd790a1e634ac5da4f108cc63111f0cdc14",
		"gitRepo":     "xgeekshq/sie-az-prototype",
		"gitProvider": "github.com",
	})
}
