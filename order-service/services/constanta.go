package services

const (
	OrderStatusPending   = "Pending"
	OrderStatusConfirmed = "Confirmed"
	OrderStatusRejected  = "Rejected"
)

func GetOrderStatus(status int64) string {
	switch status {
	case 1:
		return OrderStatusConfirmed
	case 2:
		return OrderStatusRejected
	}
	return "Unknown"
}
