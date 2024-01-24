package api

import (
	"net/http"

	"github.com/gabtec/demo-go-api/data"
	"github.com/labstack/echo/v4"
)

func GetGithubHandler(e echo.Context) error {
	v := data.ReadData("github")
	return e.JSON(http.StatusOK, v)
	// return e.JSON(http.StatusOK, map[string]string{
	// 	"version":     "v0.1.0",
	// 	"gitSHA":      "c4ab8b471cc71903757d19895e29b2bea1da3d38",
	// 	"gitRepo":     "xgeekshq/sie-az-prototype",
	// 	"gitProvider": "github.com",
	// })
}
