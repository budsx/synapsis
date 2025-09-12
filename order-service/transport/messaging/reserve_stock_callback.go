package messaging

import (
	"context"
	"encoding/json"

	"github.com/budsx/synapsis/order-service/entity"
	"github.com/budsx/synapsis/order-service/services"
)

func ReserveStockCallback(service services.OrderService) func([]byte) error {
	return func(reserveStockReq []byte) error {
		const (
			funcName = "ReserveStockCallback"
		)

		var reserveStock *entity.ReserveStockRequest
		err := json.Unmarshal(reserveStockReq, &reserveStock)
		if err != nil {
			return err
		}

		err = service.ReserveStockCallback(context.Background(), &entity.ReserveStockCallbackRequest{
			ProductID: reserveStock.ProductID,
			Quantity:  reserveStock.Quantity,
		})
		if err != nil {
			return err
		}
		return nil
	}
}
