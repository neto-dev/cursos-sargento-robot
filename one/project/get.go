package project

import "github.com/labstack/echo/v4"

func Get(){
	e:= echo.New()
	e.GET("/get", func(context echo.Context) error {
		return context.HTML(200, "<h1>Hola Mundo desde Get!</h1>")
	})

	e.GET("/get/:id", func(context echo.Context) error {
		id := context.Param("id")
		return context.HTML(200, "<h1>Hola Mundo desde Get!</h1><br><h2> Este es el id: " + id + "<h2>")
	})

	e.Start(":1234")
}
