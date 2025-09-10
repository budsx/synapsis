package services

import (
	"context"

	"github.com/budsx/synapsis/inventory-service/entity"
	"github.com/budsx/synapsis/inventory-service/repository"
)

type InventoryService interface {
	CheckStock(ctx context.Context, productID int64) (*entity.CheckStockResponse, error)
	ReserveStock(ctx context.Context, request *entity.ReserveStockRequest) (*entity.ReserveStockResponse, error)
	ReleaseStock(ctx context.Context, request *entity.ReleaseStockRequest) (*entity.ReleaseStockResponse, error)
	GetProductByID(ctx context.Context, productID int64) (*entity.Product, error)
}

type inventoryService struct {
	repo *repository.Repository
}

func NewInventoryService(repo *repository.Repository) InventoryService {
	return &inventoryService{
		repo: repo,
	}
}
