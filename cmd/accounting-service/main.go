package main

import "github.com/labstack/echo/v4"

//Main function
func main() {
	e := echo.New()
	e.GET("/", indexHandler)
}