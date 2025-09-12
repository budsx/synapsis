package services

import (
	"context"

	"github.com/budsx/synapsis/inventory-service/entity"
)

func (s *inventoryService) CheckStock(ctx context.Context, request *entity.CheckStockRequest) (*entity.CheckStockResponse, error) {
	s.logger.Info(ctx, "CheckStock", "request", request)
	stock, err := s.repo.DBReadWriter.CheckStock(ctx, request.ProductID)
	if err != nil {
		s.logger.Error(ctx, "CheckStock", "error", err)
		return nil, err
	}
	s.logger.Info(ctx, "CheckStock", "stock", stock)
	return &entity.CheckStockResponse{
		ProductID: request.ProductID,
		Stock:     stock,
	}, nil
}

func (s *inventoryService) ReserveStock(ctx context.Context, request *entity.ReserveStockRequest) error {
	s.logger.Info(ctx, "ReserveStock", "request", request)
	err := s.repo.DBReadWriter.ReserveStock(ctx, request.ProductID, request.Quantity)
	if err != nil {
		s.logger.Error(ctx, "ReserveStock", "error", err)
		return err
	}

	s.logger.Info(ctx, "ReserveStock", "success")
	return nil
}

func (s *inventoryService) ReleaseStock(ctx context.Context, request *entity.ReleaseStockRequest) error {
	s.logger.Info(ctx, "ReleaseStock", "request", request)
	err := s.repo.DBReadWriter.ReleaseStock(ctx, request.ProductID, request.Quantity)
	if err != nil {
		s.logger.Error(ctx, "ReleaseStock", "error", err)
		return err
	}
	s.logger.Info(ctx, "ReleaseStock", "success")
	return nil
}

func (s *inventoryService) GetProductByID(ctx context.Context, productID int64) (*entity.Product, error) {
	s.logger.Info(ctx, "GetProductByID", "productID", productID)
	result, err := s.repo.DBReadWriter.GetProductByID(ctx, productID)
	if err != nil {
		s.logger.Error(ctx, "GetProductByID", "error", err)
		return nil, err
	}
	s.logger.Info(ctx, "GetProductByID", "result", result)
	return result, nil
}
