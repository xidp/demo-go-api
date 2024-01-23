package main

import (
	"fmt"
	"log"

	"github.com/gabtec/demo-go-api/router"
	gou "github.com/gabtec/gabtec-gou"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Hello Universe")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := gou.GetEnv("PORT", "3000")
	addr := ":" + port

	e := echo.New()

	router.SetupRoutes(e)

	e.Start(addr)
}
