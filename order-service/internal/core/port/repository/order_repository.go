package repository

import (
	"context"

	"github.com/babaYaga451/go-zomato/order-service/internal/core/domain/entity"
)

type OrderRepository interface {
	SaveOrderAndInitiatePaymentTx(ctx context.Context, order *entity.Order) error
	FindByTrackingId(ctx context.Context, orderId string) (*entity.Order, error)
}
