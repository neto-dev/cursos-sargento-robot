package project

import "github.com/labstack/echo/v4"

func Delete(){
	e:= echo.New()
	e.DELETE("/delete", func(context echo.Context) error {
		return context.HTML(200, "<h1>Hola Mundo desde Get!</h1>")
	})

	e.DELETE("/delete/:id", func(context echo.Context) error {
		id := context.Param("id")
		params := context.QueryParam("parametro")
		return context.HTML(200, "<h1>Hola Mundo desde Get!</h1><br><h2> Este es el id: " + id + "<h2><br>" + params)
	})

	e.Start(":1234")
}
