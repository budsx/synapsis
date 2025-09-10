package repository

import (
	"github.com/budsx/synapsis/inventory-service/repository/interfaces"
	"github.com/budsx/synapsis/inventory-service/repository/postgres"
	"github.com/budsx/synapsis/inventory-service/repository/rabbitmq"
)

type Repository struct {
	DBReadWriter interfaces.InventoryDBReadWriter
	MessageQueue interfaces.MessageQueue
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
	DBConf       DBConf
	RabbitmqConf RabbitmqConf
}

func NewRepository(conf RepoConf) (*Repository, error) {
	postgresRepo, err := postgres.NewPostgresRepository(conf.DBConf.DBHost, conf.DBConf.DBPort, conf.DBConf.DBUser, conf.DBConf.DBPassword, conf.DBConf.DBName, conf.DBConf.DBDriver)
	if err != nil {
		return nil, err
	}

	rabbitmqRepo, err := rabbitmq.NewRabbitMQRepository(conf.RabbitmqConf.RabbitmqURL)
	if err != nil {
		return nil, err
	}

	return &Repository{
		DBReadWriter: postgresRepo,
		MessageQueue: rabbitmqRepo,
	}, nil
}

func (r *Repository) Close() {
	r.MessageQueue.Close()
	r.DBReadWriter.Close()
}
