package domain

import "time"

const (
	FullQualifiedEventName = "ddc_backend.datastore.event.1.created"
	DatastoreCreatedEventType = "DATASTORE_CREATED"
)

type DatastoreCreatedEvent struct {
	eventType string
	timestamp time.Time
	name      string
}

func NewDatastoreCreatedEvent(name string) *DatastoreCreatedEvent {
	return &DatastoreCreatedEvent{
		eventType: DatastoreCreatedEventType,
		timestamp: time.Now(),
		name:      name,
	}
}

func (dce *DatastoreCreatedEvent) FullQualifiedEventName() string {
	return FullQualifiedEventName
}

func (dce *DatastoreCreatedEvent) GetName() string {
	return dce.name
}

func (dce *DatastoreCreatedEvent) GetTimestamp() time.Time {
	return dce.timestamp
}

func (dce *DatastoreCreatedEvent) GetType() string {
	return DatastoreCreatedEventType
}
