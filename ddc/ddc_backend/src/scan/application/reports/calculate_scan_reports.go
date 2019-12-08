package reports

import (
	"ddc.example.com/datastore/domain"
	shared "ddc.example.com/shared/domain"
	"fmt"
)

type CalculateScanReports struct {
}

func NewCalculateScanReports() *CalculateScanReports {
	return &CalculateScanReports{}
}

func (csr *CalculateScanReports) SubscribedTo() shared.DomainEvent {
	return &domain.DatastoreCreatedEvent{}
}

func (csr *CalculateScanReports) Consume(event shared.DomainEvent) {

	fmt.Printf("Calculating scan reports ...")

}
