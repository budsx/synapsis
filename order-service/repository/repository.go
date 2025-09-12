package repository

import (
	"fmt"

	"github.com/budsx/synapsis/order-service/repository/interfaces"
	"github.com/budsx/synapsis/order-service/repository/inventoryclient"
	"github.com/budsx/synapsis/order-service/repository/postgres"
	"github.com/budsx/synapsis/order-service/repository/rabbitmq"
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
	RabbitmqURL       string
	TopicReserveStock string
	TopicReleaseStock string
}

type RepoConf struct {
	MicroConf    MicroConf
	DBConf       DBConf
	RabbitmqConf RabbitmqConf
}

type Repository struct {
	interfaces.InventoryClient
	interfaces.OrderDBReadWriter
	interfaces.MessageQueue
}

func NewRepository(conf RepoConf) (*Repository, error) {
	inventoryHost := fmt.Sprintf("%s:%d", conf.MicroConf.InventoryHost, conf.MicroConf.InventoryPort)
	inventoryClient, err := inventoryclient.NewInventoryClient(inventoryHost)
	if err != nil {
		return nil, err
	}

	orderRepo, err := postgres.NewPostgresRepository(conf.DBConf.DBHost, conf.DBConf.DBPort, conf.DBConf.DBUser, conf.DBConf.DBPassword, conf.DBConf.DBName, conf.DBConf.DBDriver)
	if err != nil {
		return nil, err
	}

	rabbitmqRepo, err := rabbitmq.NewRabbitMQRepository(conf.RabbitmqConf.RabbitmqURL, conf.RabbitmqConf.TopicReserveStock, conf.RabbitmqConf.TopicReleaseStock)
	if err != nil {
		return nil, err
	}

	return &Repository{
		InventoryClient:   inventoryClient,
		OrderDBReadWriter: orderRepo,
		MessageQueue:      rabbitmqRepo,
	}, nil
}
