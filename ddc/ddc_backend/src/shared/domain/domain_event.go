package domain

type DomainEvent interface {
	FullQualifiedEventName() string
	GetType() string
}
