package main

import (
	"ddc.example.com/datastore/application/creation"
	"github.com/labstack/echo"

	"ddc.example.com/datastore/infrastructure"
)

func main() {
	e := echo.New()

	r := infrastructure.NewDatastoreRepositoryInMemory()
	dc := creation.NewDatastoreCreation(r)
	infrastructure.NewDatastorePostController(e, dc)

	e.Logger.Fatal(e.Start(":1323"))
}
