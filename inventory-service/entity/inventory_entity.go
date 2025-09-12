package entity

type Product struct {
	ID        int64
	Name      string
	Stock     int64
	CreatedAt string
	UpdatedAt string
}

type CheckStockRequest struct {
	ProductID int64
}

type CheckStockResponse struct {
	ProductID int64
	Stock     int64
}

type ReserveStockResponse struct {
	Success bool
}

type ReleaseStockResponse struct {
	Success bool
}

type GetProductByIDResponse struct {
	Product *Product
}

type ReserveStockRequest struct {
	ProductID      int64
	Quantity       int64
	IdempotencyKey string
}

type ReleaseStockRequest struct {
	ProductID int64
	Quantity  int64
}
