package entity

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
	Status    string
	CreatedAt string
	UpdatedAt string
}

type CreateOrderRequest struct {
	ProductID      int64
	Quantity       int64
	IdempotencyKey string
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
}

type ReleaseStockCallbackRequest struct {
	ProductID int64
	Quantity  int64
}
