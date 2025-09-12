package services

import (
	"context"

	"github.com/budsx/synapsis/order-service/entity"
	"github.com/budsx/synapsis/order-service/repository"
	"github.com/budsx/synapsis/order-service/utils/common"
)

type OrderService interface {
	CreateOrder(ctx context.Context, req *entity.CreateOrderRequest) (*entity.CreateOrderResponse, error)
	ReserveStockCallback(ctx context.Context, req *entity.ReserveStockCallbackRequest) error
	ReleaseStockCallback(ctx context.Context, req *entity.ReleaseStockCallbackRequest) error
}

type orderService struct {
	repo   *repository.Repository
	logger *common.Logger
}

func NewOrderService(repo *repository.Repository, logger *common.Logger) OrderService {
	return &orderService{
		repo: repo,
	}
}
