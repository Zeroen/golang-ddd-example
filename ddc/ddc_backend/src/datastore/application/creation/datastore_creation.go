package creation

import "ddc.example.com/datastore/domain"

type DatastoreCreation struct {
}

func NewDatastoreCreation() *DatastoreCreation {
	return &DatastoreCreation{}
}

func (dc *DatastoreCreation) Invoke(id *domain.DatastoreID,
	name *domain.DatastoreName,
	ipAddress *domain.DatastoreIpAddress,
	path *domain.DatastorePath) error {

	domain.NewDatastore(id, path, ipAddress, name)

	return nil
}
