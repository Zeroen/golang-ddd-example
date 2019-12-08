package domain

import "errors"

// DatastorePath it's a value object
type DatastorePath struct {
	path string
}

func NewDatastorePath(path string) (*DatastorePath, error) {

	var err error
	dp := &DatastorePath{
		path: path,
	}

	err = dp.ensureDatastorePathCannotBeEmpty()
	if err != nil {
		return nil, err
	}

	err = dp.ensureDatastorePathShouldHavePathFormat()
	if err != nil {
		return nil, err
	}

	return dp, nil
}

func (d *DatastorePath) ensureDatastorePathCannotBeEmpty() error {
	if d.path == "" {
		return errors.New("Path of the datastore can not be empty.")
	}
	return nil
}

func (d *DatastorePath) ensureDatastorePathShouldHavePathFormat() error {
	if len(d.path) < 5 {
		return errors.New("The path should have a datastore path format.")
	}
	return nil
}

func (d *DatastorePath) Value() string {
	return d.path
}