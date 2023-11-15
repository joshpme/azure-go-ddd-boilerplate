package instruments

import (
	"ddd-boilerplate/internal/domain"
	"errors"
)

type CreateInstrumentCommand struct {
	ID               InstrumentId
	Name             string
	Type             InstrumentType
	Status           CommissionStatus
	Location         Location
	Container        Container
	CollectionPeriod int
}

func (s CreateInstrumentCommand) Partition() string {
	return string(s.ID)
}

func (s CreateInstrumentCommand) CommandHandler(existingEvents []domain.IEvent) ([]domain.IEvent, error) {
	newEvents := make([]domain.IEvent, 0)

	if err := s.Validate(); err != nil {
		return newEvents, err
	}

	if len(existingEvents) > 0 {
		return newEvents, errors.New("instrument already exists")
	}

	newEvent := CreateInstrumentEvent{
		ID:               s.ID,
		Name:             s.Name,
		Type:             s.Type,
		Status:           s.Status,
		Location:         s.Location,
		Container:        s.Container,
		CollectionPeriod: s.CollectionPeriod,
	}

	newEvents = append(newEvents, newEvent)

	return newEvents, nil
}

func (s CreateInstrumentCommand) Validate() error {
	if !s.ID.IsValid() {
		return errors.New("id must be provided and must only contain uppercase numbers and dashes")
	}

	if s.Name == "" {
		return errors.New("name must be provided")
	}

	if !s.Type.IsValid() {
		return errors.New("type must be in the list provided")
	}

	if !s.Status.IsValid() {
		return errors.New("status must be in the list provided")
	}

	if s.Location.Name == "" {
		return errors.New("location must be provided")
	}
	return nil
}

func InstrumentCommandHandler(handler domain.IDatabaseHandler, c domain.ICommand) error {
	events, err := handler.GetEvents(c.Partition())
	if err != nil {
		return errors.New("Failed to fetch events: " + err.Error())
	}

	ia := InstrumentAggregate{}
	ia = ia.Reducer(events).(InstrumentAggregate)

	newEvents, err := c.CommandHandler(events)
	if err != nil {
		return err
	}

	ia = ia.Reducer(newEvents).(InstrumentAggregate)

	err = handler.SaveEvents(c.Partition(), newEvents)
	if err != nil {
		return errors.New("Failed to save events: " + err.Error())
	}

	err = handler.SaveAggregate(ia)
	if err != nil {
		return errors.New("Failed to save aggregate: " + err.Error())
	}

	return nil
}
