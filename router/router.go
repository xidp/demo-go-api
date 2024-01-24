package router

import (
	"github.com/gabtec/demo-go-api/api"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {

	e.GET("/", api.GetHelloHandler)

	e.GET("/health", api.GetHealthHandler)

	e.GET("/version", api.GetVersionHandler)

	e.GET("/github", api.GetGithubHandler)
	e.GET("/gitlab", api.GetGitlabHandler)
}
