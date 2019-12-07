package main

import (
	"github.com/labstack/echo"

	"ddc.example.com/datastore/infrastructure"
)

func main() {
	e := echo.New()


	infrastructure.NewDatastorePostController(e)

	e.Logger.Fatal(e.Start(":1323"))
}
