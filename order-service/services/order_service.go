package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/budsx/synapsis/order-service/entity"
	inventory "github.com/budsx/synapsis/order-service/repository/inventoryclient/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *orderService) CreateOrder(ctx context.Context, req *entity.CreateOrderRequest) (*entity.CreateOrderResponse, error) {
	const (
		funcName = "CreateOrder"
	)
	s.logger.Info(fmt.Sprintf("%s %s", funcName, "Upper"), "Request", req)

	allow, err := s.repo.Redis.DeduplicateCreateOrder(ctx, req.IdempotencyKey)
	if err != nil {
		s.logger.Error(fmt.Sprintf("%s %s", funcName, "Error"), err)
		return nil, err
	}
	if !allow {
		s.logger.Info(fmt.Sprintf("%s %s", funcName, "Duplicate request"), "request", req)
		return &entity.CreateOrderResponse{
			Message: OrderStatusRejected.String(),
		}, status.Error(codes.PermissionDenied, "duplicate request")
	}

	stock, err := s.repo.InventoryClient.CheckStock(ctx, &inventory.CheckStockRequest{
		ProductId: req.ProductID,
	})
	if err != nil {
		s.logger.Error(fmt.Sprintf("%s %s", funcName, "Error"), err)
		return nil, err
	}

	if stock.Stock < req.Quantity || stock.Stock == 0 {
		s.logger.Info(fmt.Sprintf("%s %s", funcName, "Stock not enough"), "StockResponse", stock)
		return &entity.CreateOrderResponse{
			Message: OrderStatusRejected.String(),
		}, nil
	}

	// TODO: Write order to database
	_, err = s.repo.OrderDBReadWriter.CreateOrder(ctx, &entity.CreateOrderRequest{
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
		Status:    OrderStatusPending.String(),
	})
	if err != nil {
		s.logger.Error(fmt.Sprintf("%s %s", funcName, "Error"), err)
		return nil, err
	}

	err = s.repo.MessageQueue.PublishReserveStock(ctx, entity.ReserveStockRequest{
		ProductID:      req.ProductID,
		Quantity:       req.Quantity,
		IdempotencyKey: req.IdempotencyKey,
	})
	if err != nil {
		s.logger.Error(fmt.Sprintf("%s %s", funcName, "Error"), err)
		return nil, errors.New("failed to publish reserve stock")
	}

	s.logger.Info(fmt.Sprintf("%s %s", funcName, "Success"), "Request", req)
	return &entity.CreateOrderResponse{
		Message: OrderStatusPending.String(),
	}, nil
}

func (s *orderService) ReserveStockCallback(ctx context.Context, req *entity.ReserveStockCallbackRequest) error {
	const (
		funcName = "ReserveStockCallback"
	)
	s.logger.Info(fmt.Sprintf("%s %s", funcName, "Request"), "Request", req)

	// Logic:
	// if reserve status success, update order status to success
	// if reserve status failed, update order status to rejected
	if req.Status == ReserveStockStatusSuccess.ToInt32() {
		err := s.repo.OrderDBReadWriter.UpdateOrderStatus(ctx, &entity.UpdateOrderStatusRequest{
			OrderID: req.OrderID,
			Status:  OrderStatusConfirmed.String(),
		})
		if err != nil {
			s.logger.Error(fmt.Sprintf("%s %s", funcName, "Error"), err)
			return err
		}
	} else if req.Status == ReserveStockStatusFailed.ToInt32() {
		err := s.repo.OrderDBReadWriter.UpdateOrderStatus(ctx, &entity.UpdateOrderStatusRequest{
			OrderID: req.OrderID,
			Status:  OrderStatusRejected.String(),
		})
		if err != nil {
			s.logger.Error(fmt.Sprintf("%s %s", funcName, "Error"), err)
			return err
		}
	} else {
		s.logger.Warn(fmt.Sprintf("%s %s", funcName, "Invalid reserve stock status"), "Status", req.Status)
		return nil
	}

	s.logger.Info(fmt.Sprintf("%s %s", funcName, "Success"))
	return nil
}
