package main

import (
	"fmt"
	"go_enum/controllers"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Debug = true
	e.Use(middleware.CORS())
	{
		e.GET("/basic", controllers.FnBasic)
		e.GET("/old", controllers.FnOld)
		e.GET("/tip", controllers.FnIota)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
