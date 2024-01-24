package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetGitlabHandler(e echo.Context) error {
	return e.JSON(http.StatusOK, map[string]string{
		"version":     "v1.2.3",
		"gitSHA":      "1fe5fbec3d7bb7c83e598240cac3d81731e59c65",
		"gitRepo":     "ki-group-pt/xgeekshq/external/sx-frontend",
		"gitProvider": "gitlab.com",
	})
}
