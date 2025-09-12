package services

import (
	"context"
	"log/slog"

	"github.com/budsx/synapsis/order-service/entity"
	inventory "github.com/budsx/synapsis/order-service/repository/inventoryclient/proto"
)

func (s *orderService) CreateOrder(ctx context.Context, req *entity.CreateOrderRequest) (*entity.CreateOrderResponse, error) {
	slog.Info("Create Order", "request", req)
	stock, err := s.repo.InventoryClient.CheckStock(ctx, &inventory.CheckStockRequest{
		ProductId: req.ProductID,
	})
	if err != nil {
		return nil, err
	}
	slog.Info("Check Stock", "stock", stock.Stock)

	if stock.Stock < req.Quantity || stock.Stock == 0 {
		return &entity.CreateOrderResponse{
			Message: "Rejected",
		}, nil
	}

	// Publish Reserve Stock to Inventory Service
	err = s.repo.MessageQueue.PublishReserveStock(ctx, entity.ReserveStockRequest{
		ProductID:      req.ProductID,
		Quantity:       req.Quantity,
		IdempotencyKey: req.IdempotencyKey,
	})
	if err != nil {
		return nil, err
	}

	return &entity.CreateOrderResponse{
		Message: "Confirmed",
	}, nil
}

// TODO: Consumer Callback from Inventory Service
func (s *orderService) ReserveStockCallback(ctx context.Context, req *entity.ReserveStockCallbackRequest) error {
	return nil
}

// TODO: Consumer Callback from Inventory Service
func (s *orderService) ReleaseStockCallback(ctx context.Context, req *entity.ReleaseStockCallbackRequest) error {
	return nil
}
