package services

import (
	"context"

	"github.com/budsx/synapsis/inventory-service/entity"
)

func (s *inventoryService) CheckStock(ctx context.Context, request *entity.CheckStockRequest) (*entity.CheckStockResponse, error) {
	s.logger.Info("CheckStock", "request", request)
	stock, err := s.repo.DBReadWriter.CheckStock(ctx, request.ProductID)
	if err != nil {
		s.logger.Error("CheckStock", "error", err)
		return nil, err
	}
	s.logger.Info("CheckStock", "stock", stock)
	return &entity.CheckStockResponse{
		ProductID: request.ProductID,
		Stock:     stock,
	}, nil
}

func (s *inventoryService) ReserveStock(ctx context.Context, request *entity.ReserveStockRequest) error {
	s.logger.Info("ReserveStock", "request", request)
	err := s.repo.DBReadWriter.ReserveStock(ctx, request.ProductID, request.Quantity)
	if err != nil {
		s.logger.Error("ReserveStock", "error", err)
		return err
	}

	s.logger.Info("ReserveStock", "success")
	return nil
}

func (s *inventoryService) ReleaseStock(ctx context.Context, request *entity.ReleaseStockRequest) error {
	s.logger.Info("ReleaseStock", "request", request)
	err := s.repo.DBReadWriter.ReleaseStock(ctx, request.ProductID, request.Quantity)
	if err != nil {
		s.logger.Error("ReleaseStock", "error", err)
		return err
	}
	s.logger.Info("ReleaseStock", "success")
	return nil
}

func (s *inventoryService) GetProductByID(ctx context.Context, productID int64) (*entity.Product, error) {
	s.logger.Info("GetProductByID", "productID", productID)
	result, err := s.repo.DBReadWriter.GetProductByID(ctx, productID)
	if err != nil {
		s.logger.Error("GetProductByID", "error", err)
		return nil, err
	}
	s.logger.Info("GetProductByID", "result", result)
	return result, nil
}
