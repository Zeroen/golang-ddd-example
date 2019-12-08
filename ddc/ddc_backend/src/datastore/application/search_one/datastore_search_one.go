package search

import (
	"ddc.example.com/datastore/domain"
)

type DatastoreSearchOne struct {
	r domain.DatastoreRepository
}

func NewDatastoreSearchOne(r domain.DatastoreRepository) *DatastoreSearchOne {
	return &DatastoreSearchOne{
		r: r,
	}
}

func (dc *DatastoreSearchOne) Invoke(id *domain.DatastoreID) (*domain.Datastore, error) {

	d, err := dc.r.Search(id)
	if err != nil {
		return nil, err
	}

	return d, nil
}
