package messaging

import (
	"github.com/budsx/synapsis/order-service/services"
)

func ReleaseStockCallback(service services.OrderService) func([]byte) error {
	return func(paymentResponse []byte) error {
		// var releaseStock *entity.ReleaseStockRequest
		// err := json.Unmarshal(paymentResponse, &releaseStock)
		// if err != nil {
		// 	return err
		// }

		// _, err = service.ReleaseStock(context.Background(), releaseStock)
		// if err != nil {
		// 	return err
		// }

		// return nil

		return nil
	}

}
