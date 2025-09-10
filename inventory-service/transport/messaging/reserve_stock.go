package messaging

import (
	"context"
	"encoding/json"

	"github.com/budsx/synapsis/inventory-service/entity"
	"github.com/budsx/synapsis/inventory-service/services"
)

func ReserveStock(service services.InventoryService) func([]byte) error {
	return func(paymentResponse []byte) error {
		var reserveStock *entity.ReserveStockRequest
		err := json.Unmarshal(paymentResponse, &reserveStock)
		if err != nil {
			return err
		}

		_, err = service.ReserveStock(context.Background(), reserveStock)
		if err != nil {
			return err
		}

		return nil
	}
}
