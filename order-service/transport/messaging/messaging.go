package messaging

import (
	"fmt"
	"log/slog"

	"github.com/budsx/synapsis/order-service/config"
	"github.com/budsx/synapsis/order-service/repository/interfaces"
	"github.com/budsx/synapsis/order-service/services"
)

const (
	subsName = ".order.service"
)

func NewTransportOrderMessaging(conf *config.Config, service services.OrderService, messaging interfaces.MessageQueue) error {
	client := messaging.GetClient()
	if client == nil {
		return fmt.Errorf("rabbitmq client is nil")
	}

	slog.Info("Subscribing to reserve stock callback", "topic", conf.TopicReserveStockCallback)
	if err := client.Subscribe(conf.TopicReserveStockCallback, fmt.Sprintf("%s.%s", conf.TopicReserveStockCallback, subsName), ReserveStockCallback(service)); err != nil {
		return err
	}

	return nil
}
