package entity

import "time"

type ReserveStockRequest struct {
	ProductID      int64
	Quantity       int64
	IdempotencyKey string
}

type ReleaseStockRequest struct {
	ProductID int64
	Quantity  int64
}

type Order struct {
	OrderID   int64
	UserID    int64
	ProductID int64
	Quantity  int64
	Status    int32
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateOrderRequest struct {
	ProductID      int64
	Quantity       int64
	IdempotencyKey string
	Status         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type CreateOrderResponse struct {
	Message string
}

type GetOrderByIDRequest struct {
	OrderID int64
}

type ReserveStockCallbackRequest struct {
	ProductID int64
	Quantity  int64
	Status    int32
}
