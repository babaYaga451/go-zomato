package port

import (
	"context"

	"github.com/babaYaga451/go-zomato/order-service/internal/domain"
	"github.com/babaYaga451/go-zomato/order-service/internal/domain/event"
)

type CustomerRepository interface {
	FindCustomer(ctx context.Context, customerID string) (*domain.Customer, error)
}

type RestaurantRepository interface {
	FindRestaurantByProducts(ctx context.Context, restaurantID string, productIds []string) (*domain.Restaurant, error)
}

type OrderRepository interface {
	SaveOrderAndInitiatePaymentTx(ctx context.Context, order *domain.Order, orderPaymentEvent *event.OrderPaymentEvent) error
	FindByTrackingId(ctx context.Context, orderId string) (*domain.Order, error)
}
