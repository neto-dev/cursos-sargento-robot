package project

import "github.com/labstack/echo/v4"

func Put(){
	e := echo.New()
	e.PUT("/put", func(contex echo.Context) error {
		return contex.HTML(200, "<b>Hola desde Put</b>")
	})

	e.PUT("/put/:parametro", func(contex echo.Context) error {
		parametro:= contex.Param("parametro")
		queryParam := contex.QueryParam("query")
		return contex.HTML(200, "<b>Hola desde Put Con parametro: </b> " + parametro + " Query Param: " + queryParam)
	})

	e.Start(":1234")
}