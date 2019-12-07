package infrastructure

import (
	"ddc.example.com/datastore/application/creation"
	"ddc.example.com/datastore/domain"
	"github.com/labstack/echo"
	"net/http"
)

// DatastorePostController The controller used to create datastores
type DatastorePostController struct {
	dc creation.DatastoreCreation
}

// CreationDatastoreParam Body that will be receive when a user wants to create a datastores
type CreationDatastoreParam struct {
	Name string
	Ip   string
	Path string
}

// NewDatastorePostController used to instantiate the controller
func NewDatastorePostController(e *echo.Echo, dc creation.DatastoreCreation) (c *DatastorePostController) {

	// Create the DatastorePostController object
	dpc := &DatastorePostController{
		dc: dc,
	}

	// Register endpoint
	g := e.Group("/datastore")
	g.POST("/datastore/:id", dpc.Invoke)

	// Return the created object
	return dpc
}

// Invoke this method calls the controller
func (c *DatastorePostController) Invoke(ctx echo.Context) error {

	// Read the id from path parameters.
	id := ctx.QueryParam("id")

	// Read the PutDatastoreParam parameters from the body
	var param CreationDatastoreParam
	err := ctx.Bind(&param)
	if err != nil {
		return err
	}

	// Create value object for DatastoreID
	dId, err := domain.NewDatastoreId(id)
	if err != nil {
		return err
	}

	// Create value object for DatastoreName
	dName, err := domain.NewDatastoreName(param.Name)
	if err != nil {
		return err
	}

	// Create value object for DatastoreName
	dIp, err := domain.NewDatastoreIpAddress(param.Ip)
	if err != nil {
		return err
	}

	// Create value object for DatastoreName
	dPath, err := domain.NewDatastorePath(param.Path)
	if err != nil {
		return err
	}

	// Invoke the use case
	err = c.dc.Invoke(dId, dName, dIp, dPath)
	if err != nil {
		return err
	}

	// If there is no error, return an HTTP 200 Ok.
	return ctx.String(http.StatusOK, "")
}
