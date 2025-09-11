package messaging

import (
	"fmt"

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

	if err := client.Subscribe(conf.ReserveStockCallbackExchange, fmt.Sprintf("%s.%s", conf.ReserveStockCallbackExchange, subsName), ReserveStockCallback(service)); err != nil {
		return err
	}

	if err := client.Subscribe(conf.ReleaseStockCallbackExchange, fmt.Sprintf("%s.%s", conf.ReleaseStockCallbackExchange, subsName), ReleaseStockCallback(service)); err != nil {
		return err
	}

	return nil
}
