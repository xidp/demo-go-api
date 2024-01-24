package router

import (
	"github.com/gabtec/demo-go-api/api"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {

	e.GET("/", api.GetHelloHandler)

	e.GET("/health", api.GetHealthHandler)

	e.GET("/version", api.GetVersionHandler)
	e.PUT("/version", api.PutVersionHandler)

	e.GET("/github", api.GetGithubHandler)
	e.PUT("/github", api.PutGithubHandler)

	e.GET("/gitlab", api.GetGitlabHandler)
	e.PUT("/gitlab", api.PutGitlabHandler)
}
