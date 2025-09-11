package interfaces

import (
	"context"
	"io"

	"github.com/budsx/synapsis/order-service/entity"
	inventory "github.com/budsx/synapsis/order-service/repository/inventoryclient/proto"
	"github.com/budsx/synapsis/order-service/utils/common"
)

type InventoryClient interface {
	GetInventoryClient() inventory.InventoryServiceClient
	CheckStock(ctx context.Context, req *inventory.CheckStockRequest) (*inventory.CheckStockResponse, error)
}

type OrderDBReadWriter interface {
	io.Closer
	CreateOrder(ctx context.Context, req *entity.CreateOrderRequest) (*entity.CreateOrderResponse, error)
	GetOrderByID(ctx context.Context, req *entity.GetOrderByIDRequest) (*entity.Order, error)
}

type MessageQueue interface {
	io.Closer
	GetClient() *common.RabbitMQClient
	PublishReserveStock(ctx context.Context, req entity.ReserveStockRequest) error
	PublishReleaseStock(ctx context.Context, req entity.ReleaseStockRequest) error
}
