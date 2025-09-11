package messaging

import "github.com/budsx/synapsis/order-service/services"

func ReserveStockCallback(service services.OrderService) func([]byte) error {
	return func(reserveStockResponse []byte) error {
		return nil
	}
}
