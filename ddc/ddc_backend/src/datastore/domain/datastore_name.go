package domain

import "errors"

// DatastoreName it's a value object
type DatastoreName struct {
	name string
}

func NewDatastoreName(name string) (*DatastoreName, error) {

	var err error
	dn := &DatastoreName{
		name: name,
	}

	err = dn.ensureDatastoreNameCannotBeEmpty()
	if err != nil {
		return nil, err
	}

	err = dn.ensureDatastoreNameShouldHaveAtLeast5CharsOfLenght()
	if err != nil {
		return nil, err
	}

	return dn, nil
}

func (d *DatastoreName) ensureDatastoreNameCannotBeEmpty() error {
	if d.name == "" {
		return errors.New("Name of the datastore can not be empty.")
	}
	return nil
}

func (d *DatastoreName) ensureDatastoreNameShouldHaveAtLeast5CharsOfLenght() error {
	if len(d.name) >= 5 {
		return errors.New("Name of the datastore should have at least 5 characters.")
	}
	return nil
}
