package interfaces

import (
	"context"
	"io"

	"github.com/budsx/synapsis/inventory-service/entity"
	"github.com/budsx/synapsis/inventory-service/utils/common"
)

type InventoryDBReadWriter interface {
	io.Closer
	CheckStock(context.Context, int64) (int64, error)
	ReserveStock(context.Context, int64, int64) error
	ReleaseStock(context.Context, int64, int64) error
	GetProductByID(context.Context, int64) (*entity.Product, error)
}

type MessageQueue interface {
	io.Closer
	GetClient() *common.RabbitMQClient
}
