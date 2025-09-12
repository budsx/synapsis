package messaging

import (
	"context"
	"encoding/json"

	"github.com/budsx/synapsis/inventory-service/entity"
	"github.com/budsx/synapsis/inventory-service/services"
)

func ReserveStock(service services.InventoryService) func([]byte) error {
	return func(request []byte) error {
		var reserveStock *entity.ReserveStockRequest
		err := json.Unmarshal(request, &reserveStock)
		if err != nil {
			return err
		}

		err = service.ReserveStock(context.Background(), reserveStock)
		if err != nil {
			return err
		}

		return nil
	}
}
