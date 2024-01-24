package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetGitlabHandler(e echo.Context) error {
	return e.JSON(http.StatusOK, map[string]string{
		"version":     "v0.2.0",
		"gitSHA":      "6257f89aca3e5a1e30a40b4ed2b9b055fd43d85b",
		"gitRepo":     "ki-group-pt/xgeekshq/external/sx-frontend",
		"gitProvider": "gitlab.com",
	})
}
