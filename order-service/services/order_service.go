package services

import (
	"context"
	"errors"
	"log/slog"

	"github.com/budsx/synapsis/order-service/entity"
	inventory "github.com/budsx/synapsis/order-service/repository/inventoryclient/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *orderService) CreateOrder(ctx context.Context, req *entity.CreateOrderRequest) (*entity.CreateOrderResponse, error) {
	allow, err := s.repo.Redis.DeduplicateCreateOrder(ctx, req.IdempotencyKey)
	if err != nil {
		return nil, err
	}
	if !allow {
		return &entity.CreateOrderResponse{
			Message: OrderStatusRejected.String(),
		}, status.Error(codes.PermissionDenied, "duplicate request")
	}

	stock, err := s.repo.InventoryClient.CheckStock(ctx, &inventory.CheckStockRequest{
		ProductId: req.ProductID,
	})
	if err != nil {
		return nil, err
	}

	if stock.Stock < req.Quantity || stock.Stock == 0 {
		return &entity.CreateOrderResponse{
			Message: OrderStatusRejected.String(),
		}, nil
	}

	// Write order to database

	err = s.repo.MessageQueue.PublishReserveStock(ctx, entity.ReserveStockRequest{
		ProductID:      req.ProductID,
		Quantity:       req.Quantity,
		IdempotencyKey: req.IdempotencyKey,
	})
	if err != nil {
		slog.ErrorContext(ctx, "failed to publish reserve stock", "error", err)
		return nil, errors.New("failed to publish reserve stock")
	}

	return &entity.CreateOrderResponse{
		Message: OrderStatusConfirmed.String(),
	}, nil
}

func (s *orderService) ReserveStockCallback(ctx context.Context, req *entity.ReserveStockCallbackRequest) error {
	return nil
}

func (s *orderService) ReleaseStockCallback(ctx context.Context, req *entity.ReleaseStockCallbackRequest) error {
	return nil
}
