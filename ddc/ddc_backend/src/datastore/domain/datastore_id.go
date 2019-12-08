package domain

import (
	"errors"
	"regexp"
)

const UuidRegExp = "[0-9a-fA-F]{8}\\-[0-9a-fA-F]{4}\\-[0-9a-fA-F]{4}\\-[0-9a-fA-F]{4}\\-[0-9a-fA-F]{12}"

// DatastoreID it's a value object
type DatastoreID struct {
	id string
}

func NewDatastoreId(id string) (*DatastoreID, error) {

	var err error
	di := &DatastoreID{
		id: id,
	}

	err = di.ensureDatastoreIdCannotBeEmpty()
	if err != nil {
		return nil, err
	}

	err = di.ensureDatastoreIdHasUuidFormat()
	if err != nil {
		return nil, err
	}

	return di, nil
}

func (d *DatastoreID) ensureDatastoreIdCannotBeEmpty() error {
	if d.id == "" {
		return errors.New("Datastore ID of the datastore can not be empty.")
	}
	return nil
}

func (d *DatastoreID) ensureDatastoreIdHasUuidFormat() error {
	match, _ := regexp.MatchString(UuidRegExp, d.id)
	if match {
		return errors.New("ID of the datastore should be complaint with UUID format.")
	}
	return nil
}

func (d *DatastoreID) GetID() string {
	return d.id
}