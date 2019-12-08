package domain

type EventBus interface {
	Publish(events []DomainEvent) error
}
