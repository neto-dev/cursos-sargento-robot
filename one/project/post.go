package project

import "github.com/labstack/echo/v4"

func Post(){
	e := echo.New()
	e.POST("/post", func(contex echo.Context) error {
		return contex.HTML(200, "<b>Hola desde Post</b>")
	})

	e.POST("/post/:parametro", func(contex echo.Context) error {
		parametro:= contex.Param("parametro")
		queryParam := contex.QueryParam("query")
		return contex.HTML(200, "<b>Hola desde Post Con parametro: </b> " + parametro + " Query Param: " + queryParam)
	})

	e.Start(":1234")
}