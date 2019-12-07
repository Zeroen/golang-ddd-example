package domain

// Datastore it's a domain class
type Datastore struct {
	id   *DatastoreID
	name *DatastoreName
	ip   *DatastoreIpAddress
	path *DatastorePath
}

func NewDatastore(id *DatastoreID, path *DatastorePath, ip *DatastoreIpAddress, name *DatastoreName) *Datastore {
	return &Datastore{
		id:   id,
		name: name,
		ip:   ip,
		path: path,
	}
}

func (d *Datastore) DatastoreId() *DatastoreID {
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
