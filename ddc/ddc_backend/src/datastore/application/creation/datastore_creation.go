package creation

import (
	"ddc.example.com/datastore/domain"
	"errors"
)

type DatastoreCreation struct {
	r domain.DatastoreRepository
}

func NewDatastoreCreation(r domain.DatastoreRepository) *DatastoreCreation {
	return &DatastoreCreation{
		r: r,
	}
}

func (dc *DatastoreCreation) Invoke(id *domain.DatastoreID,
	name *domain.DatastoreName,
	ipAddress *domain.DatastoreIpAddress,
	path *domain.DatastorePath) error {

	err := dc.ensureDatastoreDoesntExist(id)
	if err != nil {
		return nil
	}

	d := domain.NewDatastore(id, path, ipAddress, name)

	err = dc.r.Save(d)
	if err != nil {
		return err
	}

	return nil
}

func (dc *DatastoreCreation) ensureDatastoreDoesntExist(id *domain.DatastoreID) error {

	// Look for the datastore in the database
	d, err := dc.r.Search(id)
	if err != nil {
		return err
	}

	// Check that the datastore does not exist in the databasea.
	if d != nil {
		return errors.New("The datastore [%s] already exists in the database")
	}
	return nil
}
