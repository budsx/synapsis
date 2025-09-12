package handler

import (
	"context"
	"log/slog"

	"github.com/budsx/synapsis/order-service/entity"
	order "github.com/budsx/synapsis/order-service/proto"
	"github.com/budsx/synapsis/order-service/services"
	"github.com/budsx/synapsis/order-service/utils/common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	idempotencyKey := common.GetRequestHeaderByKey(ctx, "x-idempotency-key")
	if idempotencyKey == "" {
		return nil, status.Errorf(codes.InvalidArgument, "x-idempotency-key is required")
	}
	res, err := h.service.CreateOrder(ctx, &entity.CreateOrderRequest{
		ProductID:      req.ProductId,
		Quantity:       req.Quantity,
		IdempotencyKey: idempotencyKey,
	})
	if err != nil {
		return nil, err
	}
	slog.Info("Response", "response", res)
	return &order.CreateOrderResponse{
		Message: res.Message,
	}, nil
}
