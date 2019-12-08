package domain

type DomainEventSubscriber interface {
	SubscribedTo() DomainEvent
	Consume(event DomainEvent)
}

