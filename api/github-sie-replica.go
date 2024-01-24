package api

import (
	"net/http"

	"github.com/gabtec/demo-go-api/data"
	"github.com/labstack/echo/v4"
)

var githubName = "github"

func GetGithubHandler(e echo.Context) error {
	v := data.ReadData(githubName)
	return e.JSON(http.StatusOK, v)
	// return e.JSON(http.StatusOK, map[string]string{
	// 	"version":     "v0.1.0",
	// 	"gitSHA":      "c4ab8b471cc71903757d19895e29b2bea1da3d38",
	// 	"gitRepo":     "xgeekshq/sie-az-prototype",
	// 	"gitProvider": "github.com",
	// })
}

func PutGithubHandler(e echo.Context) error {
	dto := new(bodyPayload)
	if err := e.Bind(dto); err != nil {
		return e.String(http.StatusBadRequest, "bad request")
	}

	err := data.WriteToData(githubName, dto.Version, dto.Sha)

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
