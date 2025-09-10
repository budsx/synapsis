package rabbitmq

import (
	"github.com/budsx/synapsis/inventory-service/utils/common"
)

type RabbitMQClient struct {
	client *common.RabbitMQClient
}

func NewRabbitMQRepository(rabbitmqURL string) (*RabbitMQClient, error) {
	client, err := common.NewClient(rabbitmqURL)
	if err != nil {
		return nil, err
	}
	return &RabbitMQClient{
		client: client,
	}, nil
}

func (c *RabbitMQClient) GetClient() *common.RabbitMQClient {
	if c.client != nil {
		return c.client
	}
	return nil
}

func (c *RabbitMQClient) Close() error {
	return c.client.Close()
}
