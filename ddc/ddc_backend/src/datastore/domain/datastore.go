package domain

import (
	"ddc.example.com/shared/domain"
)

// Datastore it's a domain class
type Datastore struct {
	domain.AggregateRoot
	id   *DatastoreID
	name *DatastoreName
	ip   *DatastoreIpAddress
	path *DatastorePath
}

func NewDatastore(id *DatastoreID, path *DatastorePath, ip *DatastoreIpAddress, name *DatastoreName) *Datastore {
	d := &Datastore{
		id:   id,
		name: name,
		ip:   ip,
		path: path,
	}

	event := NewDatastoreCreatedEvent(name.Value())
	d.Record(event)

	return d
}

func (d *Datastore) DatastoreID() *DatastoreID {
	return d.id
}

func (d *Datastore) DatastoreName() *DatastoreName {
	return d.name
}

func (d *Datastore) DatastoreIpAddress() *DatastoreIpAddress {
	return d.ip
}

func (d *Datastore) DatastorePath() *DatastorePath {
	return d.path
}
