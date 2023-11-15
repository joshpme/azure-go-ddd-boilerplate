package instruments

import "ddd-boilerplate/internal/domain"

type InstrumentAggregate struct {
	ID        InstrumentId
	Name      string
	Type      InstrumentType
	Status    CommissionStatus
	Location  Location
	Container Container
}

const (
	Instrument domain.AggregateType = "Instrument"
)

func (ia InstrumentAggregate) PartitionKey() string {
	return "Instruments"
}

func (ia InstrumentAggregate) AggregateName() domain.AggregateType {
	return Instrument
}

func (ia InstrumentAggregate) Reducer(events []domain.IEvent) domain.IAggregate {
	result := ia
	for _, event := range events {
		switch e := event.(type) {
		case CreateInstrumentEvent:
			result = ExecuteCreateInstrumentEvent(result, e)
		}
	}
	return result
}

func ExecuteCreateInstrumentEvent(ia InstrumentAggregate, event CreateInstrumentEvent) InstrumentAggregate {
	ia.ID = event.ID
	ia.Name = event.Name
	ia.Type = event.Type
	ia.Status = event.Status
	ia.Location = event.Location
	ia.Container = event.Container

	return ia
}
