package main

import (
	"fmt"
	"github.com/GeertJohan/go.rice" //"github.com/GeertJohan/go.rice/rice"
	"github.com/labstack/echo"
	"net/http"
)

func main() {

	e := echo.New()
	box := rice.MustFindBox("template")

	fmt.Println(box)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
