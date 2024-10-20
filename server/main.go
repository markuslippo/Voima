package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handler for the root ("/") route
func helloWorld(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Hello, World!",
	})
}

func main() {
	// Create a new Echo instance
	e := echo.New()

	// Define the root route handler for "/"
	e.GET("/", helloWorld)

	// Start the server on port 8080
	e.Logger.Fatal(e.Start("localhost:8080"))
}
