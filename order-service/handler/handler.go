package handler

import (
	"context"

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
	return h.service.CreateOrder(ctx, req)
}
