package api

import (
	"net/http"

	"github.com/gabtec/demo-go-api/version"
	"github.com/labstack/echo/v4"
)

func GetVersionHandler(e echo.Context) error {
	return e.JSON(http.StatusOK, map[string]string{
		"version":     version.Version,
		"gitSHA":      version.GitSHA,
		"gitRepo":     version.GitRepo,
		"gitProvider": version.GitProvider,
	})
}
