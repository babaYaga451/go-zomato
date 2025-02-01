package event

import (
	"time"
)

type OrderPaymentEvent struct {
	OrderID       string
	CustomerID    string
	Price         float64
	PaymentStatus string
	CreatedAt     time.Time
}

func (e *OrderPaymentEvent) GetOrderID() string {
	return e.OrderID
}

func (e *OrderPaymentEvent) GetCustomerID() string {
	return e.CustomerID
}

func (e *OrderPaymentEvent) GetPrice() float64 {
	return e.Price
}

func (e *OrderPaymentEvent) GetPaymentStatus() string {
	return e.PaymentStatus
}

func (e *OrderPaymentEvent) GetCreatedAt() time.Time {
	return e.CreatedAt
}
