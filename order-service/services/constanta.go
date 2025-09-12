package services

type OrderStatus int64

type ReserveStockStatus int64

const (
	OrderStatusPending   OrderStatus = 0
	OrderStatusConfirmed OrderStatus = 1
	OrderStatusRejected  OrderStatus = 2
	OrderStatusCanceled  OrderStatus = 3
	OrderStatusSuccess   OrderStatus = 4

	ReserveStockStatusSuccess ReserveStockStatus = 1
	ReserveStockStatusFailed  ReserveStockStatus = 2
)

func (s OrderStatus) String() string {
	switch s {
	case OrderStatusPending:
		return "Pending"
	case OrderStatusConfirmed:
		return "Confirmed"
	case OrderStatusRejected:
		return "Rejected"
	case OrderStatusCanceled:
		return "Canceled"
	case OrderStatusSuccess:
		return "Success"
	}
	return "Unknown"
}

func (s ReserveStockStatus) ToInt32() int32 {
	switch s {
	case ReserveStockStatusSuccess:
		return 1
	case ReserveStockStatusFailed:
		return 2
	}
	return -1
}
