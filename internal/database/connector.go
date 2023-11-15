package database

import (
	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
)

type CosmosHandler struct {
	ConnectionString string
}

func (manager *CosmosHandler) GetClient() (*azcosmos.Client, error) {
	return azcosmos.NewClientFromConnectionString(manager.ConnectionString, nil)
}

type IStorageHandler interface {
	GetClient() (*azcosmos.Client, error)
}

func GetInstance(connectionString string) IStorageHandler {
	return &CosmosHandler{ConnectionString: connectionString}
}
