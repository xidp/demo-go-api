package main

import (
	"fmt"

	"github.com/gabtec/demo-go-api/router"
	gou "github.com/gabtec/gabtec-gou"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Println("Hello Universe")

	err := godotenv.Load()
	if err != nil {
		fmt.Println("[INFO]: NO .env file found. Using defaults")
	}

	port := gou.GetEnv("PORT", "3000")
	addr := ":" + port

	e := echo.New()

	e.Use(middleware.CORS())
	router.SetupRoutes(e)

	e.Start(addr)
}
