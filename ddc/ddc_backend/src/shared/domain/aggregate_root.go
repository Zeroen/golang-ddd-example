package domain

type AggregateRoot struct {
	recordedDomainEvents []DomainEvent
}

func (ar *AggregateRoot) PullDomainEvents() []DomainEvent {
	list := ar.recordedDomainEvents
	ar.initialize()
	return list
}

func (ar *AggregateRoot) Record(event DomainEvent) {
	if ar.isEmpty() {
		ar.initialize()
	}
	ar.recordedDomainEvents = append(ar.recordedDomainEvents, event)
}

func (ar *AggregateRoot) initialize() {
	ar.recordedDomainEvents = make([]DomainEvent, 0)
}

func (ar *AggregateRoot) isEmpty() bool {
	return ar.recordedDomainEvents == nil
}
