package service

import (
	"context"

	"github.com/babaYaga451/go-zomato/order-service/internal/core/domain/entity"
)

type OrderService interface {
	ValidateAndInitiateOrder(ctx context.Context, order *entity.Order) error
	TrackOrder(ctx context.Context, trackingId string) (*entity.Order, error)
}
