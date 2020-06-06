package project

import "github.com/labstack/echo/v4"

func Hello(){
	e := echo.New()
	 e.GET("/", func(context echo.Context) error {
		return context.String(200, "<b>Hola Mundo! Sargento Robot</b>")
	 })

	e.Start(":1234")
}