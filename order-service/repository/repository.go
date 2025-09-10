package repository

import (
	"fmt"

	"github.com/budsx/synapsis/order-service/repository/interfaces"
	"github.com/budsx/synapsis/order-service/repository/inventoryclient"
)

type MicroConf struct {
	InventoryHost string
	InventoryPort int
}

type DBConf struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	DBDriver   string
}

type RabbitmqConf struct {
	RabbitmqURL string
}

type RepoConf struct {
	MicroConf    MicroConf
	DBConf       DBConf
	RabbitmqConf RabbitmqConf
}

type Repository struct {
	interfaces.InventoryClient
}

func NewRepository(conf RepoConf) (*Repository, error) {
	inventoryHost := fmt.Sprintf("%s:%d", conf.MicroConf.InventoryHost, conf.MicroConf.InventoryPort)
	inventoryClient, err := inventoryclient.NewInventoryClient(inventoryHost)
	if err != nil {
		return nil, err
	}
	return &Repository{
		InventoryClient: inventoryClient,
	}, nil
}
