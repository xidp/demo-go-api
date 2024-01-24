package api

import (
	"net/http"

	"github.com/gabtec/demo-go-api/data"
	"github.com/labstack/echo/v4"
)

var gitlabName = "gitlab"

func GetGitlabHandler(e echo.Context) error {
	v := data.ReadData(gitlabName)
	return e.JSON(http.StatusOK, v)
	// return e.JSON(http.StatusOK, map[string]string{
	// 	"version":     "v1.2.3",
	// 	"gitSHA":      "1fe5fbec3d7bb7c83e598240cac3d81731e59c65",
	// 	"gitRepo":     "ki-group-pt/xgeekshq/external/sx-frontend",
	// 	"gitProvider": "gitlab.com",
	// })
}

func PutGitlabHandler(e echo.Context) error {
	dto := new(bodyPayload)
	if err := e.Bind(dto); err != nil {
		return e.String(http.StatusBadRequest, "bad request")
	}

	err := data.WriteToData(gitlabName, dto.Version, dto.Sha)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Could not update version.",
		})
	}
	return e.JSON(http.StatusCreated, map[string]string{
		"ok":      "Version updated.",
		"version": dto.Version,
		"gitSHA":  dto.Sha,
	})
}
