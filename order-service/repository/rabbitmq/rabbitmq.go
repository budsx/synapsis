package rabbitmq

import (
	"github.com/budsx/synapsis/order-service/utils/common"
)

type RabbitMQClient struct {
	client                       *common.RabbitMQClient
	reserveStockCallbackExchange string
	releaseStockCallbackExchange string
}

func NewRabbitMQRepository(rabbitmqURL string, reserveStockCallbackExchange string, releaseStockCallbackExchange string) (*RabbitMQClient, error) {
	client, err := common.NewClient(rabbitmqURL)
	if err != nil {
		return nil, err
	}
	return &RabbitMQClient{
		client:                       client,
		reserveStockCallbackExchange: reserveStockCallbackExchange,
		releaseStockCallbackExchange: releaseStockCallbackExchange,
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
