package services

import (
	"context"

	"github.com/budsx/synapsis/inventory-service/entity"
)

func (s *inventoryService) CheckStock(ctx context.Context, request *entity.CheckStockRequest) (*entity.CheckStockResponse, error) {
	stock, err := s.repo.DBReadWriter.CheckStock(ctx, request.ProductID)
	if err != nil {
		return nil, err
	}

	return &entity.CheckStockResponse{
		ProductID: request.ProductID,
		Stock:     stock,
	}, nil
}

func (s *inventoryService) ReserveStock(ctx context.Context, request *entity.ReserveStockRequest) (*entity.ReserveStockResponse, error) {
	err := s.repo.DBReadWriter.ReserveStock(ctx, request.ProductID, request.Quantity)
	if err != nil {
		return nil, err
	}
	return &entity.ReserveStockResponse{
		Success: true,
	}, nil
}

func (s *inventoryService) ReleaseStock(ctx context.Context, request *entity.ReleaseStockRequest) (*entity.ReleaseStockResponse, error) {
	err := s.repo.DBReadWriter.ReleaseStock(ctx, request.ProductID, request.Quantity)
	if err != nil {
		return nil, err
	}
	return &entity.ReleaseStockResponse{
		Success: true,
	}, nil
}

func (s *inventoryService) GetProductByID(ctx context.Context, productID int64) (*entity.Product, error) {
	result, err := s.repo.DBReadWriter.GetProductByID(ctx, productID)
	if err != nil {
		return nil, err
	}
	return result, nil
}
