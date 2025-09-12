package services

import (
	"context"

	"github.com/budsx/synapsis/inventory-service/entity"
	"github.com/budsx/synapsis/inventory-service/repository"
	"github.com/budsx/synapsis/inventory-service/utils/common"
)

type InventoryService interface {
	CheckStock(ctx context.Context, request *entity.CheckStockRequest) (*entity.CheckStockResponse, error)
	ReserveStock(ctx context.Context, request *entity.ReserveStockRequest) error
	ReleaseStock(ctx context.Context, request *entity.ReleaseStockRequest) error
	GetProductByID(ctx context.Context, productID int64) (*entity.Product, error)
}

type inventoryService struct {
	repo   *repository.Repository
	logger *common.Logger
}

func NewInventoryService(repo *repository.Repository, logger *common.Logger) InventoryService {
	return &inventoryService{
		repo: repo,
	}
}
