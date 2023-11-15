package database

import (
	"context"
	"ddd-boilerplate/internal/domain"
	instrumentDomain "ddd-boilerplate/internal/domain/instruments"
	dtos "ddd-boilerplate/internal/dtos"
	"encoding/json"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	"time"
)

type DatabaseHandler struct {
	Client *azcosmos.Client
}

func (manager *DatabaseHandler) ViewContainer() (*azcosmos.ContainerClient, error) {
	return manager.Client.NewContainer("Configuration", "Aggregate")
}

func (manager *DatabaseHandler) EventContainer() (*azcosmos.ContainerClient, error) {
	return manager.Client.NewContainer("Configuration", "Events")
}

func getRows(partitionKey string, query string, containerClient *azcosmos.ContainerClient) ([][]byte, error) {
	pk := azcosmos.NewPartitionKeyString(partitionKey)
	queryPager := containerClient.NewQueryItemsPager(query, pk, nil)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	results := make([][]byte, 0)
	for queryPager.More() {
		queryResponse, err := queryPager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, item := range queryResponse.Items {
			results = append(results, item)
		}
	}
	return results, nil
}

func (manager *DatabaseHandler) SaveEvents(partitionKey string, events []domain.IEvent) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	pk := azcosmos.NewPartitionKeyString(partitionKey)
	eventContainer, err := manager.EventContainer()
	if err != nil {
		return err
	}

	batch := eventContainer.NewTransactionalBatch(pk)

	for _, event := range events {
		switch event.EventName() {
		case instrumentDomain.InstrumentCreatedEvent:
			dto := dtos.CreateInstrumentEventToDatabaseDto(event.(instrumentDomain.CreateInstrumentEvent))
			payload, err := json.Marshal(dto)
			if err != nil {
				return err
			}
			batch.CreateItem(payload, nil)
		}
	}
	response, err := eventContainer.ExecuteTransactionalBatch(ctx, batch, nil)
	if err != nil {
		return err
	}
	if !response.Success {
		return errors.New("failed to save events")
	}
	return nil
}

func (manager *DatabaseHandler) SaveAggregate(aggregate domain.IAggregate) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dto := interface{}(nil)
	switch aggregate.AggregateName() {
	case "Instrument":
		dto = dtos.InstrumentAggregateToDatabaseDto(aggregate.(instrumentDomain.InstrumentAggregate))
	}

	marshalled, err := json.Marshal(dto)
	if err != nil {
		return err
	}
	viewContainer, err := manager.ViewContainer()
	pk := azcosmos.NewPartitionKeyString(aggregate.PartitionKey())
	if err != nil {
		return err
	}
	_, err = viewContainer.UpsertItem(ctx, pk, marshalled, nil)
	return err
}

func (manager *DatabaseHandler) GetEvents(partitionKey string) ([]domain.IEvent, error) {
	eventContainer, err := manager.EventContainer()
	if err != nil {
		return nil, err
	}

	rows, err := getRows(partitionKey, "select * from c", eventContainer)
	if err != nil {
		return nil, err
	}

	events := make([]domain.IEvent, 0)
	for _, row := range rows {
		var eventDto dtos.EventDto
		err = json.Unmarshal(row, &eventDto)
		if err != nil {
			return nil, err
		}

		var eventType = domain.EventType(eventDto.EventType)
		switch eventType {
		case instrumentDomain.InstrumentCreatedEvent:
			var dto dtos.CreateInstrumentEventDto
			err = json.Unmarshal(row, &dto)
			if err != nil {
				return nil, err
			}
			createInstrumentEvent, err := dtos.CreateInstrumentEventFromDatabaseDto(dto)
			if err != nil {
				return nil, err
			}
			events = append(events, createInstrumentEvent)
		}
	}
	return events, nil
}
