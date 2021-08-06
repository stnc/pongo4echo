package main

import (
	"net/http"

	"github.com/flosch/pongo2/v4"
	"github.com/labstack/echo/v4"
	"github.com/stnc/pongo4echo"
)



//GetAllData all list
func GetAllData(e echo.Context) error {

	posts := []string{
		"Larry Ellison",
		"Carlos Slim Helu",
		"Mark Zuckerberg",
		"Amancio Ortega ",
		"Jeff Bezos",
		" Warren Buffet ",
		"Bill Gates",
		"selman tunç",
	}

	return e.Render(http.StatusOK, "templates/index.html",
		pongo2.Context{"title": "hello echo fw", "posts": posts})
}

func main() {
	e := echo.New()
	e.Renderer = pongo4echo.Renderer{Debug: true} //pongo2 v4 init
	e.Debug = true
	// http://localhost:8888/
	e.GET("/", GetAllData)
	// Start server
	e.Logger.Fatal(e.Start(":8888"))
}
