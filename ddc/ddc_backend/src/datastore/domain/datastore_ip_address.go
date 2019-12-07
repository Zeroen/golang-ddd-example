package domain

import (
	"errors"
	"regexp"
)

const IpAddressRegExp = "^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$"

// DatastoreIpAddress it's a value object
type DatastoreIpAddress struct {
	ip string
}

func NewDatastoreIpAddress(ip string) (*DatastoreIpAddress, error) {

	var err error
	di := &DatastoreIpAddress{
		ip: ip,
	}

	err = di.ensureIpAddressCannotBeEmpty()
	if err != nil {
		return nil, err
	}

	err = di.ensureDatastoreIdHasUuidFormat()
	if err != nil {
		return nil, err
	}

	return di, nil
}

func (d *DatastoreIpAddress) ensureIpAddressCannotBeEmpty() error {
	if d.ip == "" {
		return errors.New("IpAddress of the datastore can not be empty.")
	}
	return nil
}

func (d *DatastoreIpAddress) ensureDatastoreIpAddressHasIpFormat() error {
	match, _ := regexp.MatchString(IpAddressRegExp, d.ip)
	if match {
		return errors.New("IpAddress of the datastore should be complaint with Ip Address format.")
	}
	return nil
}
