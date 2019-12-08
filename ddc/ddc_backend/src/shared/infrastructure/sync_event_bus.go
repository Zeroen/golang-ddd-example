package infrastructure

import "ddc.example.com/shared/domain"

type SyncEventBus struct {
	register map[string][]domain.DomainEventSubscriber
}

func NewSyncEventBus(subscribers []domain.DomainEventSubscriber) *SyncEventBus {
	aseb := SyncEventBus{
		register: make(map[string][]domain.DomainEventSubscriber),
	}
	for _, s := range subscribers {
		aseb.registerOnEventBus(s)
	}
	return &aseb
}

func (r *SyncEventBus) registerOnEventBus(subscriber domain.DomainEventSubscriber) {
	e := subscriber.SubscribedTo()
	if r.register[e.GetType()] == nil {
		r.register[e.GetType()] = make([]domain.DomainEventSubscriber, 0)
	}
	r.register[e.GetType()] = append(r.register[e.GetType()], subscriber)
}

func (r *SyncEventBus) Publish(events []domain.DomainEvent) error {
	for _, e := range events {
		r.publish(e)
	}
	return nil
}

func (r *SyncEventBus) publish(event domain.DomainEvent) {
	asList := r.register[event.GetType()]
	if asList != nil && len(asList) > 0 {
		for _, as := range asList {
			r.eventConsumer(event, as)
		}
	}
}

func (r *SyncEventBus) eventConsumer(event domain.DomainEvent, subscriber domain.DomainEventSubscriber) {
	subscriber.Consume(event)
}
