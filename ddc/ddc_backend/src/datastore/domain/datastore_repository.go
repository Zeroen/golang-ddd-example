package domain

type DatastoreRepository interface {
	Save(datastore *Datastore) error
	Search(id *DatastoreID) (*Datastore, error)
}
