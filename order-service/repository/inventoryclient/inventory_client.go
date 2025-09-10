package inventoryclient

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/budsx/synapsis/order-service/utils/common"

	"github.com/budsx/synapsis/order-service/repository/interfaces"
	inventory "github.com/budsx/synapsis/order-service/repository/inventoryclient/proto"
)

type InventoryClient struct {
	inventory.InventoryServiceClient
}

func NewInventoryClient(url string) (interfaces.InventoryClient, error) {
	timeout := 30 * time.Second
	conn, err := common.GrpcClientConnection(url, common.SetupCircuitBreaker(timeout))
	if err != nil {
		fmt.Println("error", err)
		return nil, err
	}
	return &InventoryClient{
		InventoryServiceClient: inventory.NewInventoryServiceClient(conn),
	}, nil
}

func (c *InventoryClient) GetInventoryClient() inventory.InventoryServiceClient {
	return c.InventoryServiceClient
}

func (c *InventoryClient) CheckStock(ctx context.Context, req *inventory.CheckStockRequest) (*inventory.CheckStockResponse, error) {
	resp, err := c.InventoryServiceClient.CheckStock(ctx, req)
	slog.Info("Stock checked for product", "product_id", req.ProductId, "stock", resp.Stock)
	if err != nil {
		slog.Error("Error checking stock for product", "error", err)
		return nil, err
	}
	return resp, nil
}
