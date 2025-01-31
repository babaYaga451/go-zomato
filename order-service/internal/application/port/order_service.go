package port

import (
	"context"

	"github.com/babaYaga451/go-zomato/order-service/internal/adapter/inbound/http/dto"
)

type OrderService interface {
	CreateOrder(ctx context.Context, cmd *dto.CreateOrderCommand) (*dto.CreateOrderResponse, error)
	TrackOrder(ctx context.Context, trackingId string) (*dto.TrackOrderResponse, error)
}
