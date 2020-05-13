package project

import "github.com/labstack/echo/v4"

func Hello(){
	e := echo.New()

	e.GET("/", func(context echo.Context) error{
		return context.String(200, "<h1>Hola Mundo desde Get!</h1>")
	})

	e.Logger.Fatal(e.Start(":1234"))
}