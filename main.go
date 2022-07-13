package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Attempting to start server on port " + port)
	e.Static("/", "public")
	e.Static("/w3/Examples", "w3")
	e.Static("/teachable-machine", "teachable-machine")
	e.Static("/w4/Examples", "w4")
	e.Logger.Fatal(e.Start(":" + port))
}
