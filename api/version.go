package api

import (
	"net/http"

	"github.com/gabtec/demo-go-api/data"
	"github.com/labstack/echo/v4"
)

var appName = "app"

func GetVersionHandler(e echo.Context) error {
	v := data.ReadData(appName)

	return e.JSON(http.StatusOK, v)
	// return e.JSON(http.StatusOK, map[string]string{
	// 	"version":     version.Version,
	// 	"gitSHA":      version.GitSHA,
	// 	"gitRepo":     version.GitRepo,
	// 	"gitProvider": version.GitProvider,
	// })
}

type bodyPayload struct {
	Version string `param:"version" query:"version" form:"version" json:"version"`
	Sha     string `param:"sha" query:"sha" form:"sha" json:"sha"`
}

func PutVersionHandler(e echo.Context) error {
	dto := new(bodyPayload)
	if err := e.Bind(dto); err != nil {
		return e.String(http.StatusBadRequest, "bad request")
	}

	err := data.WriteToData(appName, dto.Version, dto.Sha)

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
