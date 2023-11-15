package domain

type ICommand interface {
	CommandHandler(existingEvents []IEvent) ([]IEvent, error)
	Partition() string
	Validate() error
}

type IAggregate interface {
	PartitionKey() string
	Reducer(events []IEvent) IAggregate
	AggregateName() AggregateType
}

type EventType string

type AggregateType string

type IEvent interface {
	EventName() EventType
}

type IDatabaseHandler interface {
	GetEvents(partitionKey string) ([]IEvent, error)
	SaveEvents(partitionKey string, items []IEvent) error
	SaveAggregate(aggregate IAggregate) error
}
