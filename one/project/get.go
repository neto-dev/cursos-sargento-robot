package project

import "github.com/labstack/echo/v4"

func Get(){
	e := echo.New()
	e.GET("/get", func(contex echo.Context) error {
		return contex.HTML(200, "<b>Hola desde Get</b>")
	})

	e.GET("/get/:parametro", func(contex echo.Context) error {
		parametro:= contex.Param("parametro")
		queryParam := contex.QueryParam("query")
		return contex.HTML(200, "<b>Hola desde Get Con parametro: </b> " + parametro + " Query Param: " + queryParam)
	})

	e.Start(":1234")
}