package services

import (
	"context"
	"log/slog"

	"github.com/budsx/synapsis/inventory-service/entity"
)

func (s *inventoryService) CheckStock(ctx context.Context, request *entity.CheckStockRequest) (*entity.CheckStockResponse, error) {
	stock, err := s.repo.DBReadWriter.CheckStock(ctx, request.ProductID)
	if err != nil {
		return nil, err
	}

	slog.Info("Stock", "stock", stock)
	return &entity.CheckStockResponse{
		ProductID: request.ProductID,
		Stock:     stock,
	}, nil
}

func (s *inventoryService) ReserveStock(ctx context.Context, request *entity.ReserveStockRequest) error {
	// TODO: Check to redis idempotency key
	slog.Info("Reserve Stock", "request", request)
	err := s.repo.DBReadWriter.ReserveStock(ctx, request.ProductID, request.Quantity)
	if err != nil {
		return err
	}

	slog.Info("Reserve Stock Success", "request", request)
	return nil
}

func (s *inventoryService) ReleaseStock(ctx context.Context, request *entity.ReleaseStockRequest) error {
	err := s.repo.DBReadWriter.ReleaseStock(ctx, request.ProductID, request.Quantity)
	if err != nil {
		return err
	}
	return nil
}

func (s *inventoryService) GetProductByID(ctx context.Context, productID int64) (*entity.Product, error) {
	result, err := s.repo.DBReadWriter.GetProductByID(ctx, productID)
	if err != nil {
		return nil, err
	}
	return result, nil
}
