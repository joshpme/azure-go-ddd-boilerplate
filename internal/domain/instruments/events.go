package instruments

import "ddd-boilerplate/internal/domain"

type CreateInstrumentEvent struct {
	ID               InstrumentId
	Name             string
	Type             InstrumentType
	Status           CommissionStatus
	Location         Location
	Container        Container
	CollectionPeriod int
}

const (
	InstrumentCreatedEvent domain.EventType = "InstrumentCreatedEvent"
)

func (s CreateInstrumentEvent) EventName() domain.EventType {
	return InstrumentCreatedEvent
}
