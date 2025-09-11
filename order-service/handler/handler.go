package handler

import (
	"context"

	"github.com/budsx/synapsis/order-service/entity"
	order "github.com/budsx/synapsis/order-service/proto"
	"github.com/budsx/synapsis/order-service/services"
)

type OrderHandler struct {
	order.UnimplementedOrderServiceServer
	service services.OrderService
}

func NewOrderHandler(service services.OrderService) *OrderHandler {
	return &OrderHandler{
		service: service,
	}
}

func (h *OrderHandler) CreateOrder(ctx context.Context, req *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	res, err := h.service.CreateOrder(ctx, &entity.CreateOrderRequest{
		ProductID:      req.ProductId,
		Quantity:       req.Quantity,
		IdempotencyKey: req.IdempotencyKey,
	})
	if err != nil {
		return nil, err
	}
	return &order.CreateOrderResponse{
		Message: res.Message,
	}, nil
}
