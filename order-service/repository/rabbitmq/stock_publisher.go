package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/budsx/synapsis/order-service/entity"
)

func (c *RabbitMQClient) PublishReserveStock(ctx context.Context, req entity.ReserveStockRequest) error {
	msg, err := json.Marshal(req)
	if err != nil {
		return err
	}
	client := c.GetClient()
	if client == nil {
		return fmt.Errorf("client is nil")
	}
	return client.Publish(c.reserveStockCallbackExchange, msg)
}

func (c *RabbitMQClient) PublishReleaseStock(ctx context.Context, req entity.ReleaseStockRequest) error {
	msg, err := json.Marshal(req)
	if err != nil {
		return err
	}
	client := c.GetClient()
	if client == nil {
		return fmt.Errorf("client is nil")
	}
	return client.Publish(c.releaseStockCallbackExchange, msg)
}
