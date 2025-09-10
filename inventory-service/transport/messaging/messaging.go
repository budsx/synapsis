package messaging

import (
	"fmt"

	"github.com/budsx/synapsis/inventory-service/config"
	"github.com/budsx/synapsis/inventory-service/repository/interfaces"
	"github.com/budsx/synapsis/inventory-service/services"
)

const (
	subsName = ".inventory.service"
)

func NewMessagingListener(conf *config.Config, service services.InventoryService, messaging interfaces.MessageQueue) error {
	client := messaging.GetClient()
	if client == nil {
		return fmt.Errorf("rabbitmq client is nil")
	}

	if err := client.Subscribe("reserve.stock", fmt.Sprintf("reserve.stock.%s", subsName), ReserveStock(service)); err != nil {
		return err
	}

	return nil
}
