package services

import (
	"context"
	"log/slog"

	order "github.com/budsx/synapsis/order-service/proto"
	"github.com/budsx/synapsis/order-service/repository"
	inventory "github.com/budsx/synapsis/order-service/repository/inventoryclient/proto"
)

type OrderService interface {
	CreateOrder(ctx context.Context, req *order.CreateOrderRequest) (*order.CreateOrderResponse, error)
}

type orderService struct {
	repo *repository.Repository
}

func NewOrderService(repo *repository.Repository) OrderService {
	return &orderService{
		repo: repo,
	}
}

func (s *orderService) CreateOrder(ctx context.Context, req *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	stock, err := s.repo.InventoryClient.CheckStock(ctx, &inventory.CheckStockRequest{
		ProductId: int64(1),
	})
	if err != nil {
		slog.Error("Error checking stock for product", "error", err)
		return nil, err
	}

	slog.Info("Inventory service is working", "stock", stock)

	return &order.CreateOrderResponse{
		OrderId: stock.Stock,
	}, nil
}
