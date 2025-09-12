package services

type OrderStatus int64

const (
	OrderStatusPending   OrderStatus = 0
	OrderStatusConfirmed OrderStatus = 1
	OrderStatusRejected  OrderStatus = 2
)

func (s OrderStatus) String() string {
	switch s {
	case OrderStatusPending:
		return "Pending"
	case OrderStatusConfirmed:
		return "Confirmed"
	case OrderStatusRejected:
		return "Rejected"
	}
	return "Unknown"
}
