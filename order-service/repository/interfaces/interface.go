package interfaces

import (
	"context"
	"io"

	"github.com/budsx/synapsis/inventory-service/utils/common"
	order "github.com/budsx/synapsis/order-service/proto"
	inventory "github.com/budsx/synapsis/order-service/repository/inventoryclient/proto"
)

type InventoryClient interface {
	GetInventoryClient() inventory.InventoryServiceClient
	CheckStock(ctx context.Context, req *inventory.CheckStockRequest) (*inventory.CheckStockResponse, error)
}

type OrderDBReadWriter interface {
	io.Closer
	CreateOrder(ctx context.Context, req *order.CreateOrderRequest) (*order.CreateOrderResponse, error)
}

type RabbitMQ interface {
	io.Closer
	GetClient() *common.RabbitMQClient
	PublishReserveStock(ctx context.Context, req model.ReserveStockRequest) error
}
