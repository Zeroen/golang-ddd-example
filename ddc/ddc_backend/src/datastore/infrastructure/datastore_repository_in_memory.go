package infrastructure

import "ddc.example.com/datastore/domain"

type DatastoreRepositoryInMemory struct {
	datastores map[string]*domain.Datastore
}

func NewDatastoreRepositoryInMemory() *DatastoreRepositoryInMemory {
	return &DatastoreRepositoryInMemory{
		datastores: make(map[string]*domain.Datastore),
	}
}

func (r *DatastoreRepositoryInMemory) Save(datastore *domain.Datastore) error {
	r.datastores[datastore.DatastoreID().GetID()] = datastore
	return nil
}

func (r *DatastoreRepositoryInMemory) Search(id *domain.DatastoreID) (*domain.Datastore, error) {
	return r.datastores[id.GetID()], nil
}
